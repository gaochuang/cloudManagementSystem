package common

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	ConfigFile = "config/config.yaml"
)

var (
	VIPER  *viper.Viper
	LOG    *zap.Logger
	CONFIG ServerConfig
)

var ()
