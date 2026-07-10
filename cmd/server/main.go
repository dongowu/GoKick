package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dongowu/gokick/internal/config"
)

func main() {
	// 解析命令行参数
	env := flag.String("env", "dev", "运行环境: dev/staging/prod")
	flag.Parse()

	cfg, err := config.Load(*env)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("✅ GoKick Scaffold 启动成功！\n")
	fmt.Printf("环境: %s\n", cfg.Env)
	fmt.Printf("服务端口: %d\n", cfg.Server.Port)
	fmt.Printf("运行模式: %s\n", cfg.Server.Mode)
	fmt.Printf("MySQL: %s@%s:%d/%s\n", cfg.MySQL.Username, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)
	fmt.Printf("Redis: %s:%d (db=%d)\n", cfg.Redis.Addr, cfg.Redis.PoolSize, cfg.Redis.DB)
	fmt.Printf("JWT过期时间: %v\n", cfg.JWT.Expire)
	fmt.Printf("日志级别: %s\n", cfg.Log.Level)

	// TODO: 初始化数据库、Redis、路由等
}
