package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"golang-stagging/core"
)

// InitConfig 初始化配置文件
func InitConfig(configPath string) {
	if configPath == "" {
		panic(fmt.Errorf("config path nil"))
		return
	}
	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		if configFileNotFoundError, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(configFileNotFoundError)
		} else {
			panic(fmt.Errorf("read config file %v error: %v", configPath, err))
		}
	}
	if err = v.Unmarshal(&core.Config); err != nil {
		fmt.Println(err)
	}
}
