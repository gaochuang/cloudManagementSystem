package common

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Mysql struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	DBName   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"Password" json:"Password" yaml:"password"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

func GormConfig() *gorm.Config {
	loglevel := logger.Error
	switch CONFIG.Mysql.LogMode {
	case "Silent":
		loglevel = logger.Silent
	case "Error":
		loglevel = logger.Error
	case "Warn":
		loglevel = logger.Warn
	case "Info":
		loglevel = logger.Info
	}

	return &gorm.Config{
		Logger:                                   logger.Default.LogMode(loglevel),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
