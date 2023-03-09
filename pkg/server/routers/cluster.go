package routers

import (
	"github.com/gin-gonic/gin"
	"kubernetes_management_system/pkg/kubernetes/cluster"
)

func InitClusterRouter(group *gin.RouterGroup) {
	k8sClusterRouter := group.Group("kubernetes")
	{
		k8sClusterRouter.POST("cluster", cluster.CreateK8SCluster)
		k8sClusterRouter.GET("cluster", cluster.ListK8sCluster)
		k8sClusterRouter.GET("cluster/config", cluster.GetK8SClusterConfig)
		k8sClusterRouter.POST("cluster/delete", cluster.DeleteK8SCluster)
	}
}
