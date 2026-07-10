package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Load 加载配置：base + env-specific + env vars override
func Load(env string) (*Config, error) {
	basePath := "config/config.yaml"
	envPath := filepath.Join("config", fmt.Sprintf("%s.yaml", env))

	// 1. 加载基础配置
	cfg, err := loadYAML(basePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("load base config: %w", err)
	}
	if cfg == nil {
		cfg = &Config{}
	}

	// 2. 加载环境覆盖配置
	if envCfg, err := loadYAML(envPath); err == nil {
		mergeConfig(cfg, envCfg)
	}

	// 3. 环境变量覆盖 (GONIO_*)
	applyEnvVars(cfg)

	cfg.Env = env
	return cfg, nil
}

func loadYAML(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// mergeConfig 合并覆盖配置到基础配置（只覆盖非零值）
func mergeConfig(base, override *Config) {
	// Server
	if override.Server.Port != 0 {
		base.Server.Port = override.Server.Port
	}
	if override.Server.Mode != "" {
		base.Server.Mode = override.Server.Mode
	}
	if override.Server.ReadTimeout != 0 {
		base.Server.ReadTimeout = override.Server.ReadTimeout
	}
	if override.Server.WriteTimeout != 0 {
		base.Server.WriteTimeout = override.Server.WriteTimeout
	}
	if override.Server.AutoMigrate {
		base.Server.AutoMigrate = true
	}

	// MySQL
	if override.MySQL.Host != "" {
		base.MySQL.Host = override.MySQL.Host
	}
	if override.MySQL.Port != 0 {
		base.MySQL.Port = override.MySQL.Port
	}
	if override.MySQL.Username != "" {
		base.MySQL.Username = override.MySQL.Username
	}
	if override.MySQL.Password != "" {
		base.MySQL.Password = override.MySQL.Password
	}
	if override.MySQL.Database != "" {
		base.MySQL.Database = override.MySQL.Database
	}
	if override.MySQL.MaxIdleConns != 0 {
		base.MySQL.MaxIdleConns = override.MySQL.MaxIdleConns
	}
	if override.MySQL.MaxOpenConns != 0 {
		base.MySQL.MaxOpenConns = override.MySQL.MaxOpenConns
	}
	if override.MySQL.ConnMaxLifetime != 0 {
		base.MySQL.ConnMaxLifetime = override.MySQL.ConnMaxLifetime
	}
	if override.MySQL.ConnMaxIdleTime != 0 {
		base.MySQL.ConnMaxIdleTime = override.MySQL.ConnMaxIdleTime
	}
	if override.MySQL.DialTimeout != 0 {
		base.MySQL.DialTimeout = override.MySQL.DialTimeout
	}
	if override.MySQL.ReadTimeout != 0 {
		base.MySQL.ReadTimeout = override.MySQL.ReadTimeout
	}
	if override.MySQL.WriteTimeout != 0 {
		base.MySQL.WriteTimeout = override.MySQL.WriteTimeout
	}
	if override.MySQL.PingTimeout != 0 {
		base.MySQL.PingTimeout = override.MySQL.PingTimeout
	}
	if override.MySQL.PrepareStmt {
		base.MySQL.PrepareStmt = true
	}
	if override.MySQL.SkipDefaultTransaction {
		base.MySQL.SkipDefaultTransaction = true
	}

	// Redis
	if override.Redis.Addr != "" {
		base.Redis.Addr = override.Redis.Addr
	}
	if override.Redis.Password != "" {
		base.Redis.Password = override.Redis.Password
	}
	if override.Redis.DB != 0 {
		base.Redis.DB = override.Redis.DB
	}
	if override.Redis.PoolSize != 0 {
		base.Redis.PoolSize = override.Redis.PoolSize
	}
	if override.Redis.MinIdleConns != 0 {
		base.Redis.MinIdleConns = override.Redis.MinIdleConns
	}
	if override.Redis.MaxIdleConns != 0 {
		base.Redis.MaxIdleConns = override.Redis.MaxIdleConns
	}
	if override.Redis.PoolTimeout != 0 {
		base.Redis.PoolTimeout = override.Redis.PoolTimeout
	}
	if override.Redis.DialTimeout != 0 {
		base.Redis.DialTimeout = override.Redis.DialTimeout
	}
	if override.Redis.ReadTimeout != 0 {
		base.Redis.ReadTimeout = override.Redis.ReadTimeout
	}
	if override.Redis.WriteTimeout != 0 {
		base.Redis.WriteTimeout = override.Redis.WriteTimeout
	}
	if override.Redis.ConnMaxIdleTime != 0 {
		base.Redis.ConnMaxIdleTime = override.Redis.ConnMaxIdleTime
	}
	if override.Redis.ConnMaxLifetime != 0 {
		base.Redis.ConnMaxLifetime = override.Redis.ConnMaxLifetime
	}
	if override.Redis.PingTimeout != 0 {
		base.Redis.PingTimeout = override.Redis.PingTimeout
	}
	if override.Redis.MaxRetries != 0 {
		base.Redis.MaxRetries = override.Redis.MaxRetries
	}
	if override.Redis.MinRetryBackoff != 0 {
		base.Redis.MinRetryBackoff = override.Redis.MinRetryBackoff
	}
	if override.Redis.MaxRetryBackoff != 0 {
		base.Redis.MaxRetryBackoff = override.Redis.MaxRetryBackoff
	}

	// MQ
	if override.MQ.Driver != "" {
		base.MQ.Driver = override.MQ.Driver
	}
	if override.MQ.ConsumerGroup != "" {
		base.MQ.ConsumerGroup = override.MQ.ConsumerGroup
	}
	if len(override.MQ.TopicConcurrency) > 0 {
		if base.MQ.TopicConcurrency == nil {
			base.MQ.TopicConcurrency = make(map[string]int)
		}
		for k, v := range override.MQ.TopicConcurrency {
			base.MQ.TopicConcurrency[k] = v
		}
	}
	if override.MQ.DefaultMaxLen != 0 {
		base.MQ.DefaultMaxLen = override.MQ.DefaultMaxLen
	}
	if len(override.MQ.TopicMaxLen) > 0 {
		if base.MQ.TopicMaxLen == nil {
			base.MQ.TopicMaxLen = make(map[string]int)
		}
		for k, v := range override.MQ.TopicMaxLen {
			base.MQ.TopicMaxLen[k] = v
		}
	}
	if override.MQ.TrimInterval != 0 {
		base.MQ.TrimInterval = override.MQ.TrimInterval
	}

	// JWT
	if override.JWT.Secret != "" {
		base.JWT.Secret = override.JWT.Secret
	}
	if override.JWT.Expire != 0 {
		base.JWT.Expire = override.JWT.Expire
	}

	// Log
	if override.Log.Mode != "" {
		base.Log.Mode = override.Log.Mode
	}
	if override.Log.Level != "" {
		base.Log.Level = override.Log.Level
	}
	if override.Log.SQLLevel != "" {
		base.Log.SQLLevel = override.Log.SQLLevel
	}
	if override.Log.FilePath != "" {
		base.Log.FilePath = override.Log.FilePath
	}
	if override.Log.MaxSize != 0 {
		base.Log.MaxSize = override.Log.MaxSize
	}
	if override.Log.MaxBackups != 0 {
		base.Log.MaxBackups = override.Log.MaxBackups
	}
	if override.Log.MaxAge != 0 {
		base.Log.MaxAge = override.Log.MaxAge
	}
	if override.Log.Compress {
		base.Log.Compress = true
	}

	// SMTP
	if override.SMTP.Host != "" {
		base.SMTP.Host = override.SMTP.Host
	}
	if override.SMTP.Port != 0 {
		base.SMTP.Port = override.SMTP.Port
	}
	if override.SMTP.Username != "" {
		base.SMTP.Username = override.SMTP.Username
	}
	if override.SMTP.Password != "" {
		base.SMTP.Password = override.SMTP.Password
	}
	if override.SMTP.From != "" {
		base.SMTP.From = override.SMTP.From
	}
}

// applyEnvVars 应用环境变量覆盖
// 格式：GONIO_SERVER_PORT=9090
func applyEnvVars(cfg *Config) {
	// Server
	if v := os.Getenv("GONIO_SERVER_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			cfg.Server.Port = port
		}
	}
	if v := os.Getenv("GONIO_SERVER_MODE"); v != "" {
		cfg.Server.Mode = v
	}
	if v := os.Getenv("GONIO_SERVER_AUTO_MIGRATE"); v != "" {
		cfg.Server.AutoMigrate = strings.ToLower(v) == "true"
	}

	// MySQL
	if v := os.Getenv("GONIO_MYSQL_HOST"); v != "" {
		cfg.MySQL.Host = v
	}
	if v := os.Getenv("GONIO_MYSQL_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			cfg.MySQL.Port = port
		}
	}
	if v := os.Getenv("GONIO_MYSQL_USERNAME"); v != "" {
		cfg.MySQL.Username = v
	}
	if v := os.Getenv("GONIO_MYSQL_PASSWORD"); v != "" {
		cfg.MySQL.Password = v
	}
	if v := os.Getenv("GONIO_MYSQL_DATABASE"); v != "" {
		cfg.MySQL.Database = v
	}

	// Redis
	if v := os.Getenv("GONIO_REDIS_ADDR"); v != "" {
		cfg.Redis.Addr = v
	}
	if v := os.Getenv("GONIO_REDIS_PASSWORD"); v != "" {
		cfg.Redis.Password = v
	}
	if v := os.Getenv("GONIO_REDIS_DB"); v != "" {
		if db, err := strconv.Atoi(v); err == nil {
			cfg.Redis.DB = db
		}
	}

	// JWT
	if v := os.Getenv("GONIO_JWT_SECRET"); v != "" {
		cfg.JWT.Secret = v
	}
	if v := os.Getenv("GONIO_JWT_EXPIRE"); v != "" {
		if sec, err := strconv.ParseInt(v, 10, 64); err == nil {
			cfg.JWT.Expire = time.Duration(sec) * time.Second
		}
	}

	// Log
	if v := os.Getenv("GONIO_LOG_MODE"); v != "" {
		cfg.Log.Mode = v
	}
	if v := os.Getenv("GONIO_LOG_LEVEL"); v != "" {
		cfg.Log.Level = v
	}
}
