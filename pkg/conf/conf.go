package conf

import (
	"dream/pkg/log"
	"github.com/spf13/viper"
)

// 初始化日志
func Init() {
	viper.SetConfigFile("./config/default.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.Fatal("初始化资源失败！")
		panic("初始化资源失败！")
	}

}
