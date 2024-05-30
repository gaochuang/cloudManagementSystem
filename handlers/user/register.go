package user

import (
	"errors"
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/models/user"
	"github.com/gaochuang/cloudManagementSystem/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	var user user.User
	err := utils.CheckParameters(ctx, &user)
	if err != nil {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(password)

	inter, err := userRegister(&user)
	if err != nil {
		common.LOG.Error(fmt.Sprintf("user: %v, register failed,", user.UserName), zap.Any("err: ", err))
	} else {
		response.ResultOk(0, inter, "register success", ctx)
	}
}

func userRegister(u *user.User) (userInter user.User, err error) {
	var user user.User
	err = common.DB.Where("username = ? ", u.UserName).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		return userInter, errors.New(fmt.Sprintf("user %v already exits", user.UserName))
	}
	err = common.DB.Create(u).Error
	return user, err
}
