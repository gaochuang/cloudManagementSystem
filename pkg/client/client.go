package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var Client kubernetes.Interface

func InitClient() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}
	Client, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	InitClient()
}

func GetClient() kubernetes.Interface {
	return Client
}
