package role

import (
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/models"
	"github.com/gaochuang/cloudManagementSystem/pkg/cms"
	"github.com/gin-gonic/gin"
)

func ListRoles(c *gin.Context) {
	keyword := c.Query("keyword")
	query := models.PageResult{}
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, c)
		return
	}

	roles, err := cms.CoreV1.Role().List(c.Request.Context(), &query, keyword)
	if err != nil {
		response.FailWithMessage(response.ERROR, "get role info failed", c)
		return

	}

	response.OkWithDetailed(roles, "get role info successful", c)
}
