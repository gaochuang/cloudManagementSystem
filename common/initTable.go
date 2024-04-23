package common

import (
	"github.com/gaochuang/cloudManagementSystem/models/cluster"
	"github.com/gaochuang/cloudManagementSystem/models/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func MysqlTable(db *gorm.DB) {
	err := db.AutoMigrate(
		user.User{},
		cluster.Cluster{},
	)
	if err != nil {
		LOG.Error("register table failed", zap.Any("err: ", err))
		os.Exit(0)
	}

	LOG.Info("register table success")
}
