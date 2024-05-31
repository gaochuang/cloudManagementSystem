package utils

import (
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
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

	//从http请求中解析数据，赋值给ptr
	if err := ctx.ShouldBindBodyWith(&ptr, binding.JSON); err != nil {
		log.Logger.LogWarn("parameter error", zap.Any("err: ", err))
		response.ResultFail(response.ParamError, "", "", ctx)
		return err
	}
	return nil
}
