package common

import (
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
)

type ServerConfig struct {
	Zap    config.Zap          `mapstructure:"zap" json:"zap" yaml:"zap"`
	System config.System       `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  config.MysqlOptions `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
