package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(ctx context.Context, user *models.User) (userData *models.User, err error)
	GetUserByUserName(ctx context.Context, username string) (us *models.User, err error)
	ChangePassword(ctx context.Context, name string, oldPassword string, newPassword string) error
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

func (u *user) GetUserByUserName(ctx context.Context, username string) (us *models.User, err error) {
	if err := u.db.Preload("Role").Where("username = ?", username).First(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

func (u *user) ChangePassword(ctx context.Context, name string, oldPassword string, newPassword string) error {
	us, err := u.GetUserByUserName(ctx, name)
	if nil != us {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(oldPassword)); err != nil {
		return errors.New("old password is incorrect")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost);
	if err != nil {
		return err
	}

	return u.db.Model(&models.User{}).Where("username = ?", name).Update("password", string(password)).Error

}
