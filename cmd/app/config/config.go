package config

// MysqlOptions 用于配置 MySQL 数据库连接
type MysqlOptions struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	DBName   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	UserName string `mapstructure:"username" json:"username" yaml:"user-name"`
	Password string `mapstructure:"Password" json:"Password" yaml:"password"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap   string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

// Zap 用于配置 zap 日志库的
type Zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
	Director     string `mapstructure:"director" json:"director" yaml:"director"`
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	LinkName     string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	LogInConsole bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
	EncodeLevel  string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
}

// System 用于配置系统级别的
type System struct {
	AutoMigrateDb bool `mapstructure:"auto-migrate-db" json:"autoMigrateDb" yaml:"auto-migrate-db"` //自动创建表
}

// LogOptions 用于配置通用日志
type LogOptions struct {
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	DirectorPath string `mapstructure:"directorPath" json:"director-path" yaml:"director-path"`
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
}

type HttpOptions struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type Config struct {
	Mysql  MysqlOptions `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap          `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System       `mapstructure:"system" json:"system" yaml:"system"`
	Log    LogOptions   `mapstructure:"log" json:"log" yaml:"log"`
	Http   HttpOptions  `mapstructure:"http" json:"http" yaml:"http"`
}
