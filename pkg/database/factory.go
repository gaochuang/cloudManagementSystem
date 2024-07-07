package database

import (
	"github.com/gaochuang/cloudManagementSystem/pkg/database/system"
	"github.com/gaochuang/cloudManagementSystem/pkg/database/user"
	"gorm.io/gorm"
)

type ShareFactory interface {
	User() user.UsersInterface
	System() system.SystemsSettingInterface
}

type shareFactory struct {
	db *gorm.DB
}

func NewFactory(db *gorm.DB) ShareFactory {
	return &shareFactory{
		db: db,
	}
}

func (s *shareFactory) User() user.UsersInterface {
	return user.NewUser(s.db)
}

func (s *shareFactory) System() system.SystemsSettingInterface {
	return system.NewSystem(s.db)
}
