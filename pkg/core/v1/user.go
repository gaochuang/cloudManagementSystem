package v1

import (
	"context"
	"errors"
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"github.com/gaochuang/cloudManagementSystem/internal"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/database"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//参考client-go中工厂设计模式
//https://github.com/kubernetes/client-go/blob/master/kubernetes/typed/core/v1/core_client.go

type UserGetter interface {
	User() UsersInterface
}

type UsersInterface interface {
	Create(ctx context.Context, user *models.User) (userData *models.User, err error)
	GetJwt(ctx context.Context) []byte
	GetUserByName(ctx context.Context, userName string) (userDate *models.User, err error)
	GetUserNameByContext(ctx *gin.Context) string
	ReleaseToken(ctx context.Context, user *models.User, jwtKey []byte) (token string, err error)
	ParseToken(ctx context.Context, token string, jwtKey []byte) (*jwt.Token, *internal.JwtCustomClaims, error)
	ChangePassword(ctx context.Context, name string, oldPassword string, newPassword string) error
	GetUsers(ctx context.Context) (userList []*models.UsersListResponse, err error)
	DeleteUsers(ctx context.Context, request models.DeleteUsersRequest) error
	Update(ctx context.Context, id uint, user *models.User) error
}

type user struct {
	config  config.Config
	factory database.ShareFactory
}

func newUser(p *platform) UsersInterface {
	return &user{
		config:  p.config,
		factory: p.factory,
	}
}

func (u *user) Create(ctx context.Context, user *models.User) (userData *models.User, err error) {
	if user.UserName == "" || user.Password == "" {
		log.Logger.LogWarn("user name or password is empty")
		return nil, errors.New("user name or password is empty")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.LogError("genera password failed", zap.Any("err: ", err))
		return nil, err
	}
	user.Password = string(hashPassword)
	if userData, err = u.factory.User().Create(ctx, user); err != nil {
		log.Logger.LogErrorWithCtx(ctx, "user register failed", zap.String("user name", user.UserName), zap.Error(err))
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
		log.Logger.LogErrorWithCtx(ctx, "get user error", zap.Any("user name", userName))
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
			Subject:   "user token",
		},
	}
	jwtKeyToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	if token, err = jwtKeyToken.SignedString(jwtKey); err != nil {
		log.Logger.LogErrorWithCtx(ctx, "Failed to issue the user token", zap.Error(err))
		return "", err
	}
	return token, nil
}

func (u *user) ParseToken(ctx context.Context, token string, jwtKey []byte) (*jwt.Token, *internal.JwtCustomClaims, error) {
	claims := &internal.JwtCustomClaims{}
	tokenKey, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tokenKey.Valid {
		return nil, nil, err
	}
	return tokenKey, claims, nil
}

func (u *user) GetUserNameByContext(ctx *gin.Context) string {
	if value, exists := ctx.Get("user"); !exists {
		return ""
	} else {
		user := value.(*models.User)
		return user.UserName
	}
}

func (u *user) ChangePassword(ctx context.Context, name string, oldPassword string, newPassword string) error {
	return u.factory.User().ChangePassword(ctx, name, oldPassword, newPassword)
}

func (u *user) GetUsers(ctx context.Context) (userList []*models.UsersListResponse, err error) {
	return u.factory.User().GetUsers(ctx)
}

func (u *user) DeleteUsers(ctx context.Context, request models.DeleteUsersRequest) error {
	return u.factory.User().DeleteUsers(ctx, request)
}

func (u *user) Update(ctx context.Context, id uint, user *models.User) error {
	return u.factory.User().Update(ctx, id, user)
}
