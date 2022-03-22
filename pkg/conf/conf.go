package conf

import (
	"github.com/spf13/viper"
)

// Init 初始化日志
func Init() {
	viper.SetConfigFile("./config/default.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("初始化资源失败！")
	}

}
