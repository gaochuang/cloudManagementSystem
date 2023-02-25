package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"errcode": 0, "data": gin.H{"user": user}})
}
