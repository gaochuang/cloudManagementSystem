package user

import (
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
