package config

type MysqlOptions struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	DBName   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"Password" json:"Password" yaml:"password"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap   string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

type Zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
	Director     string `mapstructure:"director" json:"director" yaml:"director"`
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	LinkName     string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	LogInConsole bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
	EncodeLevel  string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
}

type System struct {
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`
	AutoMigrateDb bool   `mapstructure:"auto-migrate-db" json:"autoMigrateDb" yaml:"auto-migrate-db"`
}

type Config struct {
	Mysql  MysqlOptions `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap          `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System       `mapstructure:"system" json:"system" yaml:"system"`
}
