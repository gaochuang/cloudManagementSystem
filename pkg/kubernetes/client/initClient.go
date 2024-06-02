package client

import (
	"errors"
	"github.com/gaochuang/cloudManagementSystem/api/response"
	"github.com/gaochuang/cloudManagementSystem/pkg/log"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/service/cluster"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strconv"
)

func CreateK8SClient(config string) (*kubernetes.Clientset, error) {

	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(config))
	if err != nil {
		log.Logger.LogError("kubeConfig content error ", zap.Any("err:", err))
		return nil, errors.New("kubeConfig content error")
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Logger.LogError("create kubernetes client failed ", zap.Any("err:", err))
		return nil, errors.New("kubeConfig content error")
	}

	return clientSet, nil
}
func GetClusterClient(ctx *gin.Context) (*kubernetes.Clientset, error) {
	clusterId := ctx.DefaultQuery("clusterId", "1")
	id, err := strconv.ParseUint(clusterId, 10, 32)
	if err != nil {
		log.Logger.LogError("cluster id is invalid ", zap.Any("err:", err))
		response.FailWithMessage(response.CreateK8SClusterError, "cluster id is invalid", ctx)
		return nil, err
	}
	cluster, err := cluster.GetCluster(uint(id))
	if err != nil {
		log.Logger.LogError("get cluster failed ", zap.Any("err:", err))
		return nil, err
	}
	client, _ := CreateK8SClient(cluster.KubeConfig)
	return client, err
}
