package conf

import (
	"github.com/toolkits/pkg/file"
	"github.com/toolkits/pkg/runner"
	"path"
)

func GetYamlFile(module string) string {
	configDir := path.Join(runner.Cwd, "config")

	path := path.Join(configDir, module+".yaml")
	if file.IsExist(path) {
		return path
	}

	return ""
}
