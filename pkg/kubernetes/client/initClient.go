package client

import (
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/response"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/service/cluster"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strconv"
)

func CreateK8SClient(config string) (*kubernetes.Clientset, error) {

	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))
	if err != nil {
		common.LOG.Error(fmt.Sprintf("kubeConfig content error, %v", zap.Any("err:", err)))
		return nil, errors.New("kubeConfig content error")
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		common.LOG.Error(fmt.Sprintf("create kubernetes client failed, %v", zap.Any("err:", err)))
		return nil, errors.New("kubeConfig content error")
	}

	return clientSet, nil
}
func GetClusterClient(ctx *gin.Context) (*kubernetes.Clientset, error) {
	clusterId := ctx.DefaultQuery("clusterId", "1")
	id, err := strconv.ParseUint(clusterId, 10, 32)
	if err != nil {
		common.LOG.Error("cluster id is invalid")
		response.FailWithMessage(response.CreateK8SClusterError, "cluster id is invalid", ctx)
		return nil, err
	}
	cluster, err := cluster.GetCluster(uint(id))
	if err != nil {
		common.LOG.Error("get cluster failed", zap.Any("err", err))
		return nil, err
	}
	client, _ := CreateK8SClient(cluster.KubeConfig)
	return client, err
}
