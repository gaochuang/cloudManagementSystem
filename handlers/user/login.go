package user

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/gaochuang/cloudManagementSystem/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var loginUser models.LoginUser
	if err := utils.CheckParameters(ctx, &loginUser); err != nil {
		return
	}

	if loginUser.UserName == "" {
		response.FailWithMessage(response.UserNameEmpty, response.ParamErrorMsg, ctx)
		return
	}

	if loginUser.Password == "" {
		response.FailWithMessage(response.UserPassEmpty, response.UserPasswordIsEmptyMsg, ctx)
		return
	}

	u, err := authentication(&loginUser)
	if err != nil {
		log.Logger.LogError("user login failed", zap.Any("err: ", err))
		response.FailWithMessage(response.InternalServerError, response.LoginCheckErrorMsg, ctx)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginUser.Password)); err != nil {
		response.FailWithMessage(response.AuthError, "", ctx)
		return
	}

	token, err := common.ReleaseToken(u)
	if err != nil {
		log.Logger.LogError("token generate failed", zap.Any("err: ", err))
		response.FailWithMessage(response.InternalServerError, fmt.Sprintf("token genetate err: %v", err), ctx)
	}

	response.OkWithDetailed(gin.H{"token": token, "username": u.UserName}, "login success", ctx)

}

func authentication(loginUser *models.LoginUser) (models.User, error) {
	var user models.User
	err := common.DB.Where("userName = ? ", loginUser.UserName).First(&user).Error
	return user, err
}
