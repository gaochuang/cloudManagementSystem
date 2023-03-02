package client

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubernetes_management_system/common"
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
