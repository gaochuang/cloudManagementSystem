package common

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ConfigFile = "config/config.yaml"
)

var (
	VIPER  *viper.Viper
	LOG    *zap.Logger
	CONFIG ServerConfig
	DB     *gorm.DB
)
