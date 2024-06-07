package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

func CheckParameters(ctx *gin.Context, ptr interface{}) error {
	if nil == ptr {
		return nil
	}

	switch v := reflect.ValueOf(ptr).Elem(); v.Kind() {
	case reflect.String:
		if v.String() != "" {
			return errors.New(v.String())
		}
	case reflect.Interface:
		if err, ok := v.Interface().(error); ok {
			return errors.New(err.Error())
		}
	}

	//从http请求中解析数据，赋值给ptr
	if err := ctx.ShouldBindJSON(&ptr); err != nil {
		return err
	}
	return nil
}
