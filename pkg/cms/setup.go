package cms

import (
	"github.com/gaochuang/cloudManagementSystem/cmd/app/options"
	"github.com/gaochuang/cloudManagementSystem/pkg/core/v1"
)

var CoreV1 v1.CoreV1Interface

func Setup(o *options.Options) {
	CoreV1 = v1.New(*o.Config)
}
