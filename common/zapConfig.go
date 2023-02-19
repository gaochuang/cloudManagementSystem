package common

type Zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
	Director     string `mapstructure:"director" json:"director" yaml:"director"`
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	LinkName     string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	LogInConsole bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
	EncodeLevel  string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
}
