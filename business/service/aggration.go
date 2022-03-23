package service

import (
	"dream/pkg/log"
	"github.com/fatih/color"
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
	result, err := NewExportExcel().Export(contents)
	if err != nil {
		log.Logger.Error(err.Error())
		color.Red("导出Excel异常：", err.Error())
		return
	}

	// todo 5.0 异常情况的考虑
	// todo 控制台打印
	d := color.New(color.FgWhite, color.Bold)
	d.Printf("恭喜您导出成功~ \n")

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("生成的Excel文件path:", result.ExcelPath)
}
