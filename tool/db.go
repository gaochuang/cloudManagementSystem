package tool

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kubernetes_management_system/common"
	"os"
)

func GormMysql() *gorm.DB {

	config := common.CONFIG.Mysql
	//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Addr + ")/" + config.DBName + "?" + config.Config

	mySqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(mySqlConfig), common.GormConfig())
	if err != nil {
		common.LOG.Error("mysql connection failed", zap.Any("err:", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			common.LOG.Error("connect DB pool failed", zap.Any("err:", err))
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(10)
		return db
	}
}
