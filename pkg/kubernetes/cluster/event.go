package cluster

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/pkg/kubernetes/client"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/service/cluster"
	"github.com/gaochuang/cloudManagementSystem/utils"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	namespace := utils.ParseNamespaceParameter(ctx)
	clusterClient, err := client.GetClusterClient(ctx)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), ctx)
		return
	}
	field := fmt.Sprintf("type=%s", "Warning")
	events, err := cluster.GetNodeEvent(clusterClient, namespace, field)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), ctx)
	}
	response.OkWithData(events, ctx)
}
