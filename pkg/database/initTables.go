package database

import (
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func MySqlTables(db *gorm.DB) error {
	//数据库 表的初始化
	err := db.AutoMigrate(
		models.User{},
		models.Cluster{},
	)
	if err != nil {
		log.Logger.LogError("create database tablse failed", zap.Any("err: ", err))
		return err
	}
	return nil
}
