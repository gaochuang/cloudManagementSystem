package database

import (
	"github.com/gaochuang/cloudManagementSystem/pkg/database/user"
	"gorm.io/gorm"
)

type ShareFactory interface {
	User() user.UserInterface
}

type shareFactory struct {
	db *gorm.DB
}

func NewFactory(db *gorm.DB) ShareFactory {
	return &shareFactory{
		db: db,
	}
}

func (s *shareFactory) User() user.UserInterface {
	return user.NewUser(s.db)
}
