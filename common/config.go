package common

import (
	"bytes"
	"github.com/gaochuang/cloudManagementSystem/pkg/conf"
	"fmt"
	"github.com/spf13/viper"
	"github.com/toolkits/pkg/file"
)

func ParseConfigFile() error {
	//parse server config file
	yamlFile := conf.GetYamlFile("server")
	if yamlFile == "" {
		return fmt.Errorf("the config file for server don't exit")
	}

	bs, err := file.ReadBytes(yamlFile)
	if err != nil {
		return fmt.Errorf("read yaml %s failed, reason: %v", yamlFile, err)
	}

	//set config file type is yaml
	viper.SetConfigType("yaml")

	err = viper.ReadConfig(bytes.NewBuffer(bs))
	if err != nil {
		return fmt.Errorf("viper read configure from %s yaml failed, reason: %v", yamlFile, err)
	}
	return nil
}
