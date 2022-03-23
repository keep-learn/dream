package service

import (
	"dream/business/dto"
	"dream/pkg/log"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"time"
)

// ExportExcel 导出输出
type ExportExcel struct {
}

func NewExportExcel() ExportExcel {
	return ExportExcel{}
}

// Export 导出数据
func (ee ExportExcel) Export(contents []dto.ExportItem) (result dto.ExportResult, err error) {

	f := excelize.NewFile()
	// 标题样式
	headerStyle, err := f.NewStyle(`{"font":{"bold":true,"family":"宋体","size":11}}`)
	if err != nil {
		log.Logger.Error("创建标题样式失败")
		err = errors.New("创建标题样式失败")
		return
	}
	// 正文样式
	bodyStyle, err := f.NewStyle(`{"font":{"family":"宋体","size":11}}`)
	if err != nil {
		log.Logger.Error("创建正文样式失败")
		err = errors.New("创建正文样式失败")
		return
	}

	// 创建一个sheet
	sheetName := "Sheet1"
	sheetIndex := f.NewSheet(sheetName)
	f.SetActiveSheet(sheetIndex)

	// 设置标题内容
	f.SetCellValue(sheetName, "A1", "等级")
	f.SetCellValue(sheetName, "B1", "序号")
	f.SetCellValue(sheetName, "C1", "类型")
	f.SetCellValue(sheetName, "D1", "优先级")
	f.SetCellValue(sheetName, "E1", "说明")

	// 设置标题内容样式
	f.SetCellStyle(sheetName, "A1", "E1", headerStyle)

	// 设置正文内容
	index := 2
	for _, item := range contents {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", index), item.Rank)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", index), item.Seq)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", index), item.Type)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", index), item.Priority)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", index), item.Desc)
		index++
	}
	// 设置正文内容样式
	f.SetCellStyle(sheetName, "A2", fmt.Sprintf("E%d", index), bodyStyle)

	// 生成Excel；暂不考虑文件名重复的情况（可以新增随机数）
	fullPath := viper.GetString("excel.path") + time.Now().Format("220060102150405") + ".xlsx"
	if err := f.SaveAs(fullPath); err != nil {
		log.Logger.Error("报错文件失败", zap.Error(err))
	}
	result = dto.ExportResult{
		ExcelPath: fullPath,
	}
	return
}
