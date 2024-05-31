package middleware

import (
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		if "" != token {
			decodeToken(token, ctx)
		} else {
			tokenString := ctx.GetHeader("token")
			if "" == tokenString || strings.HasPrefix(tokenString, "jwt") {
				ctx.JSON(http.StatusUnauthorized, gin.H{"errcode": 401, "errmsg": "not logged in or illegally accessed"})
				ctx.Abort()
				return
			}
			//tokenString = tokenString[4:] //jwt:
			decodeToken(tokenString, ctx)
		}
	}
}

func decodeToken(token string, ctx *gin.Context) {
	tk, claims, err := common.ParseToken(token)
	if err != nil || !tk.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errcode": 401, "errmsg": "authorization has expired"})
		ctx.Abort()
		return
	}

	var u models.User

	err = common.DB.Where("userName = ? ", claims.Username).First(&u).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errcode": 401, "errmsg": "authentication failed"})
	}

	ctx.Set("user", u)
	ctx.Set("claims", claims)
	ctx.Next()
}

func authentication(loginUser *models.LoginUser) (models.User, error) {
	var user models.User
	err := common.DB.Where("userName = ? ", loginUser.UserName).First(&user).Error
	return user, err
}
