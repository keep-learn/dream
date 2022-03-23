package service

import (
	"context"
	"dream/pkg/log"
	"github.com/fatih/color"
	"go.uber.org/zap"
)

// RunExportExcel  执行导出Excel的命令
func RunExportExcel(ctx context.Context, args []string) {
	// 1.0 校验用户输入
	log.New().WithContext(ctx).Info("start RunExportExcel")
	userInputService := NewInput(args)
	err := userInputService.Check(ctx)
	if err != nil {
		log.New().WithContext(ctx).Error(err.Error())
		return
	}

	// 4.0 导入excel
	contents := userInputService.Construct()
	result, err := NewExportExcel().Export(ctx, contents)
	if err != nil {
		log.New().WithContext(ctx).Error(err.Error())
		color.Red("导出Excel异常：", err.Error())
		return
	}

	// todo 5.0 异常情况的考虑
	// todo 控制台打印
	d := color.New(color.FgWhite, color.Bold)
	d.Printf("恭喜您导出成功~ \n")

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("生成的Excel文件path:", result.ExcelPath)
	log.New().WithContext(ctx).Info("end RunExportExcel", zap.String("excel_path", result.ExcelPath))
}
