package user

import (
	"github.com/gaochuang/cloudManagementSystem/models/cluster"
	"github.com/gaochuang/cloudManagementSystem/models/user"
	"gorm.io/gorm"
)

func MySqlTables(db *gorm.DB) {
	//数据库 表的初始化
	err := db.AutoMigrate(
		user.User{},
		cluster.Cluster{},
	)
	if err != nil {
	}
}
