package user

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/models/user"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

	token, err := common.ReleaseToken(u)
	if err != nil {
		common.LOG.Error("token generate failed ", zap.Any("err: ", err))
		response.FailWithMessage(response.InternalServerError, fmt.Sprintf("token genetate err: %v", err), ctx)
	}

	response.OkWithDetailed(gin.H{"token": token, "username": u.UserName}, "login success", ctx)

}

func authentication(loginUser *user.LoginUser) (user.User, error) {
	var user user.User
	err := common.DB.Where("userName = ? ", loginUser.UserName).First(&user).Error
	return user, err
}
