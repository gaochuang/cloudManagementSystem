package config

type MysqlOptions struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	DBName   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"Password" json:"Password" yaml:"password"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
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
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
