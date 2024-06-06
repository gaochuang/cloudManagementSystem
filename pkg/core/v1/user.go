package v1

import (
	"context"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"github.com/gaochuang/cloudManagementSystem/internal"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/database"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//参考client-go中工厂设计模式
//https://github.com/kubernetes/client-go/blob/master/kubernetes/typed/core/v1/core_client.go

type UserGetter interface {
	User() UserInterface
}

type UserInterface interface {
	Create(ctx context.Context, user *models.User) (userData *models.User, err error)
	GetJwt(ctx context.Context) []byte
	GetUserByName(ctx context.Context, userName string) (userDate *models.User, err error)
	ReleaseToken(ctx context.Context, user *models.User, jwtKey []byte) (token string, err error)
}

type user struct {
	config  config.Config
	factory database.ShareFactory
}

func newUser(p *platform) UserInterface {
	return &user{
		config:  p.config,
		factory: p.factory,
	}
}

func (u *user) Create(ctx context.Context, user *models.User) (userData *models.User, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.LogError("genera password failed", zap.Any("err: ", err))
		return nil, err
	}
	user.Password = string(password)
	if userData, err = u.factory.User().Create(ctx, user); err != nil {
		log.Logger.LogDebugWithCtx(ctx, "user register failed", zap.String("user name", user.UserName), zap.Error(err))
		return nil, err
	}
	return userData, nil
}

func (u *user) GetJwt(ctx context.Context) []byte {
	jwt := u.config.Http.Jwt
	if len(jwt) == 0 {
		jwt = "platform"
	}
	return []byte(jwt)
}

func (u *user) GetUserByName(ctx context.Context, userName string) (userDate *models.User, err error) {
	if userDate, err = u.factory.User().GetUserByUserName(ctx, userName); err != nil {
		log.Logger.LogErrorWithCtx(ctx, "get user errot", zap.Any("user name", userName))
		return nil, err
	}
	return userDate, nil
}

func (u *user) ReleaseToken(ctx context.Context, user *models.User, jwtKey []byte) (token string, err error) {
	expirationTime := time.Now().Add(internal.JwtTokenExpired)
	claims := &internal.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.UserName,
		UserRole: user.Role.Code,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "platform",
			Subject:   "user tokern",
		},
	}
	jwtKeyToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	if token, err = jwtKeyToken.SignedString(jwtKey); err != nil {
		log.Logger.LogErrorWithCtx(ctx, "Failed to issue the user token", zap.Error(err))
		return "", err
	}
	return token, nil
}
