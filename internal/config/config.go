package config

import "time"

type Config struct {
	Server        ServerConfig
	MySQL         MySQLConfig
	Redis         RedisConfig
	MQ            MQConfig
	JWT           JWTConfig
	Log           LogConfig
	SMTP          SMTPConfig
	Env           string `json:"env"`
}

type ServerConfig struct {
	Port         int           `json:"port"`
	Mode         string        `json:"mode"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	AutoMigrate  bool          `json:"auto_migrate"`
}

type MySQLConfig struct {
	Host              string        `json:"host"`
	Port              int           `json:"port"`
	Username          string        `json:"username"`
	Password          string        `json:"password"`
	Database          string        `json:"database"`
	MaxIdleConns      int           `json:"max_idle_conns"`
	MaxOpenConns      int           `json:"max_open_conns"`
	ConnMaxLifetime   time.Duration `json:"conn_max_lifetime"`
	ConnMaxIdleTime   time.Duration `json:"conn_max_idle_time"`
	DialTimeout       time.Duration `json:"dial_timeout"`
	ReadTimeout       time.Duration `json:"read_timeout"`
	WriteTimeout      time.Duration `json:"write_timeout"`
	PingTimeout       time.Duration `json:"ping_timeout"`
	PrepareStmt       bool          `json:"prepare_stmt"`
	SkipDefaultTransaction bool `json:"skip_default_transaction"`
}

type RedisConfig struct {
	Addr            string        `json:"addr"`
	Password        string        `json:"password"`
	DB              int           `json:"db"`
	PoolSize        int           `json:"pool_size"`
	MinIdleConns    int           `json:"min_idle_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	PoolTimeout     time.Duration `json:"pool_timeout"`
	DialTimeout     time.Duration `json:"dial_timeout"`
	ReadTimeout     time.Duration `json:"read_timeout"`
	WriteTimeout    time.Duration `json:"write_timeout"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	PingTimeout     time.Duration `json:"ping_timeout"`
	MaxRetries      int           `json:"max_retries"`
	MinRetryBackoff time.Duration `json:"min_retry_backoff"`
	MaxRetryBackoff time.Duration `json:"max_retry_backoff"`
}

type MQConfig struct {
	Driver          string            `json:"driver"`
	ConsumerGroup   string            `json:"consumer_group"`
	TopicConcurrency map[string]int   `json:"topic_concurrency"`
	DefaultMaxLen   int               `json:"default_max_len"`
	TopicMaxLen     map[string]int    `json:"topic_max_len"`
	TrimInterval   time.Duration     `json:"trim_interval"`
}

type JWTConfig struct {
	Secret string        `json:"secret"`
	Expire time.Duration `json:"expire"`
}

type LogConfig struct {
	Mode      string `json:"mode"`
	Level     string `json:"level"`
	SQLLevel  string `json:"sql_level"`
	FilePath  string `json:"file_path"`
	MaxSize   int    `json:"max_size"`
	MaxBackups int   `json:"max_backups"`
	MaxAge    int    `json:"max_age"`
	Compress  bool   `json:"compress"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	From     string `json:"from"`
	TLS      bool   `json:"tls"`
}


