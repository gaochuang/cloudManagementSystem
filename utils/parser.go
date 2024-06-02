package utils

import "github.com/gin-gonic/gin"

func ParseNamespaceParameter(request *gin.Context) string {
	return request.Query("namespace")
}
