package main

import (
	"dream/business/dto"
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
	flag.StringVar(&exportContent, "export", "", "请输入需要导出的内容！")
}

func main() {

	// 1.0 解析命令行输入
	flag.Parse()

	// 2.0 校验用户输入
	userInput := service.NewInput(exportContent)
	err := userInput.Check()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	// 3.0 初始化 日志、配置资源
	log.Init()
	conf.Init()

	// todo 4.0 导入excel
	contents := make([]dto.ExportItem, 0)
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
