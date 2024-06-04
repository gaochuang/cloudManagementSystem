package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/models"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(ctx context.Context, user *models.User) (userData *models.User, err error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserInterface {
	return &user{
		db: db,
	}
}

func (u *user) Create(ctx context.Context, user *models.User) (us *models.User, err error) {
	if !errors.Is(u.db.Where("username = ? ", user.UserName).First(&us).Error, gorm.ErrRecordNotFound) {
		return us, errors.New(fmt.Sprintf("user %v already exists", user.UserName))
	}

	err = u.db.Create(&user).Error
	return us, err
}
