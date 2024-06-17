package user

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/cms"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/gaochuang/cloudManagementSystem/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var user models.User

	err := utils.CheckParameters(ctx, &user)
	if err != nil {
		log.Logger.LogError("check parameters failed")
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, ctx)
		return
	}

	if data, err := cms.CoreV1.User().Create(ctx.Request.Context(), &user); err != nil {
		log.Logger.LogError("register failed: ", zap.Any("user: ", user.UserName), zap.Any("err: ", err))
		response.FailWithMessage(response.UserRegisterFail, err.Error(), ctx)
	} else {
		response.ResultOk(0, data, "register success", ctx)
	}
}

func Login(ctx *gin.Context) {
	var loginUser models.LoginUser
	if err := utils.CheckParameters(ctx, &loginUser); err != nil {
		log.Logger.LogError("check parameters failed")
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, ctx)
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		log.Logger.LogWarn("compare password failed: ", zap.Any("err: ", err.Error()))
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

func GetUsers(ctx *gin.Context) {
	useList, err := cms.CoreV1.User().GetUsers(ctx.Request.Context())
	if err != nil {
		response.Fail(ctx)
		return
	}

	response.OkWithData(useList, ctx)
}

func ChangePassword(ctx *gin.Context) {
	var cp models.UsersChangePasswordRequest

	if err := utils.CheckParameters(ctx, &cp); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, ctx)
		return
	}
	name := cms.CoreV1.User().GetUserNameByContext(ctx)
	if err := cms.CoreV1.User().ChangePassword(ctx.Request.Context(), name, cp.OldPassword, cp.NewPassword); err != nil {
		response.FailWithMessage(response.ERROR, err.Error(), ctx)
	}
	response.Ok(ctx)
}

func DeleteUsers(ctx *gin.Context) {
	var userIds models.DeleteUsersRequest

	if err := utils.CheckParameters(ctx, &userIds); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, ctx)
		return
	}

	if err := cms.CoreV1.User().DeleteUsers(ctx.Request.Context(), userIds); err != nil {
		response.FailWithMessage(response.ERROR, err.Error(), ctx)
		return
	}

	response.Ok(ctx)

}
