package conf

import (
	"github.com/spf13/viper"
)

// CfgFile 配置文件路径,允许在初始化前,由外部包赋值
var CfgFile string

// Init 初始化日志
func Init() {
	viper.SetConfigFile(CfgFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic("初始化资源失败！")
	}

}
