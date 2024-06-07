package middleware

import (
	"github.com/gaochuang/cloudManagementSystem/pkg/cms"
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
				ctx.JSON(http.StatusUnauthorized, gin.H{"error code": 401, "error message": "not logged in or illegally accessed"})
				ctx.Abort()
				return
			}
			decodeToken(tokenString, ctx)
		}
	}
}

func decodeToken(token string, ctx *gin.Context) {
	tk, claims, err := cms.CoreV1.User().ParseToken(ctx, token, cms.CoreV1.User().GetJwt(ctx))
	if err != nil || !tk.Valid {
		ctx.JSON(http.StatusAccepted, gin.H{"error code": 400, "error message": "authorization has expired"})
		ctx.Abort()
		return
	}

	user, err := cms.CoreV1.User().GetUserByName(ctx, claims.Username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error code": 403, "error message": "authentication failed"})
		ctx.Abort()
		return
	}

	if !*user.Status {
		ctx.JSON(http.StatusOK, gin.H{"error core": 403, "error message": "user disabled"})
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
	ctx.Set("username", user.UserName)
	ctx.Set("claims", claims)
	ctx.Next()
}
