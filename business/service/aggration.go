package service

import (
	"context"
	"dream/pkg/log"
	"github.com/fatih/color"
	"go.uber.org/zap"
)

// RunExportExcel  执行导出Excel的命令
func RunExportExcel(ctx context.Context, args []string) {

	log.New().WithContext(ctx).Info("start RunExportExcel")

	// 1.0 校验用户输入
	userInputService := NewInput(args)
	err := userInputService.Check(ctx)
	if err != nil {
		log.New().WithContext(ctx).Warn(err.Error())
		color.Red("参数异常：%s", err.Error())
		return
	}
	// 2.0 导出excel
	contents := userInputService.Construct()
	result, err := NewExportExcel().Export(ctx, contents)
	if err != nil {
		log.New().WithContext(ctx).Error(err.Error())
		color.Red("导出Excel异常：%s", err.Error())
		return
	}
	// 3.0 控制台输出结果
	color.New(color.FgWhite, color.Bold).Printf("恭喜您导出成功，生成的Excel文件path:")
	color.New(color.FgCyan).Add(color.Underline).Println(result.ExcelPath)

	log.New().WithContext(ctx).Info("end RunExportExcel", zap.String("excel_path", result.ExcelPath))
}
