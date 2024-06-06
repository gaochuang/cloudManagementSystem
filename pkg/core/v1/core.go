package v1

import (
	"github.com/gaochuang/cloudManagementSystem/cmd/app/config"
	"github.com/gaochuang/cloudManagementSystem/pkg/database"
	"k8s.io/client-go/kubernetes"
)

type CoreV1Interface interface {
	UserGetter
}

type platform struct {
	config     config.Config
	factory    database.ShareFactory
	clientSets map[string]*kubernetes.Clientset
}

func New(config config.Config, factory database.ShareFactory) CoreV1Interface {
	return &platform{
		config:  config,
		factory: factory,
	}
}

func (p *platform) User() UserInterface {
	return newUser(p)
}
