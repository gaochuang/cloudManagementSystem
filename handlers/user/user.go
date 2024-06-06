package user

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/cms"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/gaochuang/cloudManagementSystem/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(ctx *gin.Context) {
	var user models.User
	err := utils.CheckParameters(ctx, &user)
	if err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, ctx)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(password)
	if data, err := cms.CoreV1.User().Create(ctx.Request.Context(), &user); err != nil {
		log.Logger.LogError("register failed: ", zap.Any("user: ", user.UserName), zap.Any("err: ", err))
	} else {
		response.ResultOk(0, data, "register success", ctx)
	}
}

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

	user, err := cms.CoreV1.User().GetUserByName(ctx, loginUser.UserName)
	if err != nil {
		response.FailWithMessage(response.ERROR, err.Error(), ctx)
		return
	}
	if *user.Status == false {
		response.FailWithMessage(response.UserDisable, response.UserDisableMsg, ctx)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		response.FailWithMessage(response.AuthError, response.LoginCheckErrorMsg, ctx)
		return
	}

	jwtKey := cms.CoreV1.User().GetJwt(ctx)

	token, err := cms.CoreV1.User().ReleaseToken(ctx.Request.Context(), user, jwtKey)
	if err != nil {
		log.Logger.LogError("token generate failed", zap.Any("err: ", err))
		response.FailWithMessage(response.InternalServerError, fmt.Sprintf("token genetate err: %v", err), ctx)
	}

	response.OkWithDetailed(gin.H{"token": token, "username": user.UserName, "role": user.Role}, "login success", ctx)
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"errcode": 0, "data": gin.H{"user": user}})
}

func authentication(loginUser *models.LoginUser) (models.User, error) {
	var user models.User
	err := common.DB.Where("userName = ? ", loginUser.UserName).First(&user).Error
	return user, err
}
