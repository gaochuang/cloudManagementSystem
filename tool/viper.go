package tool

import (
	"github.com/gaochuang/cloudManagementSystem/common"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		config = common.ConfigFile
		fmt.Printf("get config file from default path: %s \n", config)
	} else {
		config = path[0]
		fmt.Printf("get config file from: %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config file failed"))
	}
	v.WatchConfig()
	v.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("config file changed:", event.Name)
		if err = v.Unmarshal(&common.CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&common.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
