package common

type Mysql struct {
	addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	DBName   string `mapstructure:"DBName" json:"DBName" yaml:"DBName"`
	UserName string `mapstructure:"UserName" json:"userName" yaml:"userName"`
	Password string `mapstructure:"Password" json:"Password" yaml:"Password"`
}
