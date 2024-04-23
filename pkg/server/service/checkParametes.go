package service

import (
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
