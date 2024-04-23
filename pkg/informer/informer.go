package informer

import (
	"github.com/gaochuang/cloudManagementSystem/pkg/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"time"
)

var (
	gvrs = []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "pods"},
	}
)

func NewInformer(stop <-chan struct{}) {

	client := client.GetClient()
	shareInformer := informers.NewSharedInformerFactory(client, time.Second*60)

	for _, gvr := range gvrs {
		_, err := shareInformer.ForResource(gvr)
		if err != nil {
			panic(err.Error())
		}
	}

	shareInformer.WaitForCacheSync(stop)
	shareInformer.Start(stop)
}
