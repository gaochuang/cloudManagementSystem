package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"kubernetes_management_system/common"
	"kubernetes_management_system/pkg/server/response"
)

func CheckParameters(ctx *gin.Context, ptr interface{}) error {
	if nil == ptr {
		return nil
	}

	switch i := ptr.(type) {
	case string:
		if i != "" {
			panic(i)
		}
	case error:
		panic(i.Error())
	}

	if err := ctx.ShouldBindBodyWith(&ptr, binding.JSON); err != nil {
		common.LOG.Warn(fmt.Sprintf("parameter error: %v", err.Error()))
		response.ResultFail(response.ParamError, "", "", ctx)
		return err
	}
	return nil
}
