package cluster

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubernetes_management_system/common"
	"kubernetes_management_system/pkg/kubernetes/client"
	"kubernetes_management_system/pkg/server/response"
	"kubernetes_management_system/pkg/server/service/cluster"
)

func GetEvents(ctx *gin.Context) {
	namespace := common.ParseNamespaceParameter(ctx)
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
