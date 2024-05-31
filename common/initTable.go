package common

import (
	"github.com/gaochuang/cloudManagementSystem/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func MysqlTable(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
		models.Cluster{},
	)
	if err != nil {
		LOG.Error("register table failed", zap.Any("err: ", err))
		os.Exit(0)
	}

	LOG.Info("register table success")
}
