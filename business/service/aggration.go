package service

import (
	"dream/pkg/log"
	"fmt"
)

// RunExportExcel  执行导出Excel的命令
func RunExportExcel(args []string) {
	// 1.0 校验用户输入
	userInputService := NewInput(args)
	err := userInputService.Check()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	// 4.0 导入excel
	contents := userInputService.Construct()
	err = NewExportExcel().Export(contents)
	if err != nil {
		// todo 日志
		return
	}

	// todo 5.0 异常情况的考虑
	// todo 控制台打印
	fmt.Println("success")
}
