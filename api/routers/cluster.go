package routers

import (
	"github.com/gaochuang/cloudManagementSystem/pkg/kubernetes/cluster"
	"github.com/gin-gonic/gin"
)

func InitClusterRouter(group *gin.RouterGroup) {
	k8sClusterRouter := group.Group("kubernetes")
	{
		k8sClusterRouter.POST("cluster", cluster.CreateK8SCluster)
		k8sClusterRouter.GET("cluster", cluster.ListK8sCluster)
		k8sClusterRouter.GET("cluster/config", cluster.GetK8SClusterConfig)
		k8sClusterRouter.POST("cluster/delete", cluster.DeleteK8SCluster)
		k8sClusterRouter.GET("cluster/info", cluster.GetK8SClusterDetailInfo)
		k8sClusterRouter.GET("event", cluster.GetEvents)
	}
}
