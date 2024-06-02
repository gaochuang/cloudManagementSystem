package v1

import (
	"context"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

//参考client-go中工厂设计模式
//https://github.com/kubernetes/client-go/blob/master/kubernetes/typed/core/v1/core_client.go

type UserGetter interface {
	User() UserInterface
}

type UserInterface interface {
	Create(ctx context.Context, user *models.User) (userData *models.User, err error)
}

type user struct {
	config config.Config
}

func newUser(p *platform) UserInterface {
	return &user{
		config: p.config,
	}
}

func (u *user) Create(ctx context.Context, user *models.User) (userData *models.User, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.LogError("genera password failed", zap.Any("err: ", err))
		return nil, err
	}
	user.Password = string(password)
	return
}
