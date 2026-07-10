# GoKick Scaffold

一个生产级、易用、高性能的 Go 后端脚手架。

## 特性

- 🔧 **多环境配置**：支持 dev/staging/prod 配置分离 + 环境变量覆盖
- 🛡️ **统一错误处理**：结构化错误码 + 自动 HTTP 状态映射
- ⚡ **Redis 限流**：基于 Lua 脚本的 IP+路由+方法三级限流
- 🔐 **JWT 认证**：双角色（app/admin）支持
- 📨 **消息队列**：Watermill + Redis Streams / MySQL
- 📊 **可观测性**：OpenTelemetry + Prometheus + Jaeger（开发中）
- 📚 **API 文档**：Swagger/OpenAPI 自动生成（开发中）
- 🐳 **云原生**：Docker + Compose + Kubernetes（开发中）

## 快速开始

### 环境要求

- Go 1.26+
- MySQL 8.0+
- Redis 6.0+

### 安装

```bash
git clone https://github.com/dongowu/gokick.git
cd gokick
go mod tidy
```

### 配置

复制环境配置模板：

```bash
cp config/dev.yaml.example config/dev.yaml
# 编辑 config/dev.yaml 填入你的数据库配置
```

### 运行

```bash
# 开发环境（默认）
go run cmd/server/main.go

# 指定环境
GONIO_ENV=prod go run cmd/server/main.go

# 环境变量覆盖
GONIO_SERVER_PORT=9090 go run cmd/server/main.go
```

### CLI 代码生成器

```bash
# 生成 User 模块的完整代码
go run cmd/gen/main.go generate User --module=user

# 生成 Order 模块
go run cmd/gen/main.go generate Order --module=order
```

## 项目结构

```
gokick/
├── cmd/
│   ├── server/          # 服务入口
│   └── gen/             # CLI 代码生成器
├── config/              # 配置文件
│   ├── config.yaml      # 基础配置
│   ├── dev.yaml         # 开发环境
│   ├── staging.yaml     # 预发布环境
│   └── prod.yaml        # 生产环境
├── internal/
│   ├── config/          # 配置加载器
│   ├── handler/         # HTTP 处理器
│   ├── service/         # 业务逻辑层
│   ├── repository/      # 数据访问层
│   ├── model/           # 数据模型
│   ├── middleware/      # 中间件
│   ├── pkg/             # 通用工具包
│   │   ├── apperror/    # 统一错误处理
│   │   ├── response/    # 统一响应格式
│   │   ├── validator/   # 参数校验
│   │   └── ratelimit/   # Redis Lua 限流
│   ├── mq/              # 消息队列
│   ├── database/        # 数据库连接
│   ├── cache/           # Redis 连接
│   ├── svc/             # ServiceContext 依赖注入
│   └── telemetry/       # OpenTelemetry 可观测性
├── docs/                # 项目文档
├── scripts/             # 脚本工具
├── deploy/              # 部署配置
└── examples/            # 示例项目
```

## 核心中间件

### 限流

基于 Redis Lua 脚本的原子限流，支持 IP + 路由 + 方法三维度：

```go
// 配置示例
ratelimit.RegisterRule("GET", "/api/v1/products", 1, time.Second) // 1请求/秒
ratelimit.RegisterRule("POST", "/api/v1/orders", 1, 3*time.Second) // 1请求/3秒
```

### JWT 认证

```go
// 生成 Token
token, err := middleware.GenerateToken(userID, "app")

// 验证 Token
claims, err := middleware.ParseToken(token)
```

### 统一错误处理

```go
// 返回业务错误
c.Error(apperror.ErrInvalidParams)

// 包装普通错误
c.Error(apperror.Wrap(err, 1005, "创建失败"))
```

## 贡献

欢迎提交 Issue 和 PR。

## 许可证

MIT
