package main

import (
	"dream/business/service"
	"dream/pkg/conf"
	"dream/pkg/log"
	"flag"
	"fmt"
)

// 用户输入
var exportContent string

func init() {
	// 获取用户输入
	flag.StringVar(&exportContent, "export", "X 信息、Y 信息；甲类、乙类、丁类", "请输入需要导出的内容！")
}

func main() {

	// 1.0 初始化 日志、配置资源
	log.Init()
	conf.Init()

	// 2.0 解析命令行输入
	flag.Parse()

	// 3.0 校验用户输入
	userInputService := service.NewInput(exportContent)
	err := userInputService.Check()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	// todo 4.0 导入excel
	contents := userInputService.Construct()
	err = service.NewExportExcel().Export(contents)
	if err != nil {
		// todo 日志
		return
	}

	// todo 5.0 异常情况的考虑
	// todo 控制台打印
	// todo 校验参数合法性
	fmt.Println("inputContent:", exportContent)
}
