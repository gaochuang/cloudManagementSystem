package routers

import (
	"github.com/gin-gonic/gin"
	"kubernetes_management_system/pkg/kubernetes/cluster"
)

func InitClusterRouter(group *gin.RouterGroup) {
	k8sClusterRouter := group.Group("kubernetes")
	{
		k8sClusterRouter.POST("cluster", cluster.CreateK8SCluster)
	}
}
