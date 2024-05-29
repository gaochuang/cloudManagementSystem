package cluster

import (
	"fmt"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/models/cluster"
	"github.com/gaochuang/cloudManagementSystem/pkg/kubernetes/client"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/service"
	cluster2 "github.com/gaochuang/cloudManagementSystem/pkg/server/service/cluster"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreateK8SCluster(c *gin.Context) {

	var cluster cluster.Cluster

	if err := service.CheckParameters(c, &cluster); err != nil {
		common.LOG.Error("check cluster parameters fai`led")
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

func ListK8sCluster(c *gin.Context) {
	query := cluster.PaginationQ{}

	//can be queried only
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, c)
		return
	}

	var clusters []cluster.Cluster

	if err := cluster2.ListCluster(&query, &clusters); err != nil {
		common.LOG.Error("get cluster failed", zap.Any("err", err))
		response.FailWithMessage(response.InternalServerError, "get cluster failed", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			Data:  clusters,
			Total: query.Total,
			Size:  query.Size,
			Page:  query.Page,
		}, "get cluster successful", c)
	}

}

func GetK8SClusterConfig(c *gin.Context) {
	clusterId := c.DefaultQuery("clusterId", "1")
	id, err := strconv.ParseUint(clusterId, 10, 32)
	if err != nil {
		common.LOG.Error("cluster id is invalid", zap.Any("err", err))
		response.FailWithMessage(response.InternalServerError, "cluster id is invalid", c)
		return
	}

	clusterConfig, err := cluster2.GetCluster(uint(id))
	if err != nil {
		common.LOG.Error("get cluster config failed", zap.Any("err", err))
		response.FailWithMessage(response.InternalServerError, "get cluster config failed", c)
		return
	}
	data := map[string]interface{}{"config": clusterConfig.KubeConfig, "clusterName": clusterConfig.ClusterName}
	response.OkWithData(data, c)
}

func DeleteK8SCluster(c *gin.Context) {
	var id cluster.ClusterIds
	if err := service.CheckParameters(c, &id); err != nil {
		return
	}
	if err := cluster2.DeleteCluster(id); err != nil {
		userName, _ := c.Get("username")
		common.LOG.Error(fmt.Sprintf("user: %s, delete cluster failed", userName))
		response.FailWithMessage(response.InternalServerError, "delete cluster failed", c)
		return
	}
	response.Ok(c)
	return
}

func GetK8SClusterDetailInfo(c *gin.Context) {
	clusterClient, err := client.GetClusterClient(c)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}

	info, err := cluster2.GetClusterInfo(clusterClient)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}
	response.OkWithData(info, c)
}
