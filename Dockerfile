# 构建阶段
FROM golang:1.26-alpine AS builder

# 安装构建依赖
RUN apk add --no-cache git make ca-certificates

WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/gokick cmd/server/main.go

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/gokick .

# 复制配置文件（可选，生产中通过环境变量注入）
COPY --from=builder /app/config ./config

# 创建日志目录
RUN mkdir -p logs

EXPOSE 8080

CMD ["./gokick", "--env=prod"]
