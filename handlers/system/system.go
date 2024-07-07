package system

import (
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/cms"
	"github.com/gaochuang/cloudManagementSystem/utils"
	"github.com/gin-gonic/gin"
)

func SetSystemSafe(ctx *gin.Context) {
	var systemSafeRequest models.SystemSafeSettingsRequest

	if err := utils.CheckParameters(ctx, &systemSafeRequest); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, ctx)

		return
	}

	if err := cms.CoreV1.SystemSetting().SystemSafeSettings(ctx, &systemSafeRequest); err != nil {
		response.FailWithMessage(response.ERROR, err.Error(), ctx)
		return
	}
	response.Ok(ctx)

}

func GetSystemSafe(ctx *gin.Context) {

	settings, err := cms.CoreV1.SystemSetting().GetSystemSafeSettings(ctx.Request.Context())
	if err != nil {
		response.FailWithMessage(response.ERROR, err.Error(), ctx)
	}

	response.OkWithData(settings, ctx)
}
