package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cores() gin.HandlerFunc {
	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		//OPTIONS will return 204 no content
		if "OPTIONS" == context.Request.Method {
			context.AbortWithStatus(http.StatusNoContent)
		}
		//do request
		context.Next()
	}
}
