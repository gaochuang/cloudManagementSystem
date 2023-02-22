package common

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"kubernetes_management_system/models/user"
	"os"
)

func MysqlTable(db *gorm.DB) {

	err := db.AutoMigrate(
		user.User{},
	)
	if err != nil {
		LOG.Error("register table failed", zap.Any("err: ", err))
		os.Exit(0)
	}

	LOG.Info("register table success")
}
