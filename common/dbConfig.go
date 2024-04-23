package common

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
