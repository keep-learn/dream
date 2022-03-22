package conf

import (
	"dream/pkg/log"
	"github.com/spf13/viper"
)

// 初始化日志
func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.Fatal("初始化资源失败！")
		panic("初始化资源失败！")
	}

}
