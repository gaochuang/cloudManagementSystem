package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"kubernetes_management_system/common"
	"kubernetes_management_system/models/user"
	"kubernetes_management_system/pkg/server/response"
	"kubernetes_management_system/pkg/server/service"
)

func Login(ctx *gin.Context) {
	var loginUser user.LoginUser

	if err := service.CheckParameters(ctx, &loginUser); err != nil {
		return
	}

	if loginUser.UserName == "" {
		response.FailWithMessage(response.UserNameEmpty, "", ctx)
		return
	}

	if loginUser.Password == "" {
		response.FailWithMessage(response.UserPassEmpty, "", ctx)
		return
	}

	u, err := authentication(&loginUser)
	if err != nil {
		common.LOG.Error("user login failed", zap.Any("err: ", err))
		response.FailWithMessage(response.InternalServerError, "", ctx)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginUser.Password)); err != nil {
		response.FailWithMessage(response.AuthError, "", ctx)
		return
	}

	response.Ok(ctx)
}

func authentication(loginUser *user.LoginUser) (user.User, error) {
	var user user.User
	err := common.DB.Where("userName = ? ", loginUser.UserName).First(&user).Error
	return user, err
}
