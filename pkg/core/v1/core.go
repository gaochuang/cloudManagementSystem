package v1

import (
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"k8s.io/client-go/kubernetes"
)

type CoreV1Interface interface {
	UserGetter
}

type platform struct {
	config     config.Config
	clientSets map[string]*kubernetes.Clientset
}

func New(config config.Config) CoreV1Interface {
	return &platform{
		config: config,
	}
}

func (p *platform) User() UserInterface {
	return newUser(p)
}
