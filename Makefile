.PHONY: run run-dev run-prod gen test test-cover lint fmt clean docker-up docker-down

# 运行服务
run:
	go run cmd/server/main.go --env=dev

run-dev:
	go run cmd/server/main.go --env=dev

run-prod:
	go run cmd/server/main.go --env=prod

# 代码生成
gen:
	go run cmd/gen/main.go $(ENTITY) --module=$(MODULE)

# 测试
test:
	go test ./... -v

test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# 代码质量
lint:
	golangci-lint run

fmt:
	goimports -w .

# Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# 清理
clean:
	go clean -cache
	rm -rf coverage.out coverage.html
	rm -rf bin/
