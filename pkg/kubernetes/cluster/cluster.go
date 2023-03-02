package cluster

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kubernetes_management_system/common"
	"kubernetes_management_system/models/cluster"
	"kubernetes_management_system/pkg/kubernetes/client"
	"kubernetes_management_system/pkg/server/response"
	"kubernetes_management_system/pkg/server/service"
	cluster2 "kubernetes_management_system/pkg/server/service/cluster"
)

func CreateK8SCluster(c *gin.Context) {

	var cluster cluster.Cluster

	if err := service.CheckParameters(c, &cluster); err != nil {
		common.LOG.Error("check cluster parameters failed")
		return
	}

	client, err := client.CreateK8SClient(cluster.KubeConfig)
	if err != nil {
		response.FailWithMessage(response.CreateK8SClusterError, err.Error(), c)
		return
	}
	version, err := GetClusterVersion(client)
	if err != nil {
		response.FailWithMessage(response.CreateK8SClusterError, "connect cluster failed, please check out network", c)
		return
	}
	cluster.ClusterVersion = version
	count, err := GetClusterNodeCount(client)
	if err != nil {
		response.FailWithMessage(response.CreateK8SClusterError, "get cluster node count failed", c)
		common.LOG.Error("get cluster node count failed", zap.Any("err", err))
		return
	}
	cluster.NodeCount = count

	if err = cluster2.StorageCluster(cluster); err != nil {
		common.LOG.Error("storage cluster to db failed", zap.Any("err", err))
		response.FailWithMessage(response.CreateK8SClusterError, "storage cluster to db failed", c)
		return
	} else {
		response.OkWithMessage("create cluster successful", c)
	}

}
