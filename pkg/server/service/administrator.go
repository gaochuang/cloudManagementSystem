package service

import "github.com/gin-gonic/gin"

func checkParameters(ctx *gin.Context, ptr interface{}) error {
	if ptr == nil {
		return nil
	}
	t := ptr.(type)
	switch t {
	case string:
		if t != nil {
			t
		}
	}

}

func Login(context *gin.Context) {

	var user LoginUser

}
