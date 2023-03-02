package cluster

import (
	"context"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes_management_system/common"
)

func GetClusterVersion(client *kubernetes.Clientset) (string, error) {
	version, err := client.ServerVersion()
	if err != nil {
		common.LOG.Error("get version from cluster failed", zap.Any("err: ", err))
		return "", err
	}
	return version.String(), nil
}

func GetClusterNodeCount(client *kubernetes.Clientset) (int, error) {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(nodes.Items), nil
}
