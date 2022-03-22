package service

import (
	"dream/business/dto"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

// 导出输出
type ExportExcel struct {
}

func NewExportExcel() ExportExcel {
	return ExportExcel{}
}

// 导出数据
func (ee ExportExcel) Export(contents []dto.ExportItem) (err error) {
	f := excelize.NewFile()

	// 标题样式
	headerStyle, err := f.NewStyle(`{"font":{"bold":true,"family":"宋体","size":11}}`)
	if err != nil {
		// todo 日志
		return errors.New("创建标题样式失败")
	}
	// 正文样式
	bodyStyle, err := f.NewStyle(`{"font":{"family":"宋体","size":11}}`)
	if err != nil {
		// todo 日志
		return errors.New("创建正文样式失败")
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
	fileName := time.Now().Format("220060102150405") + ".xlsx"
	if err := f.SaveAs("./" + fileName); err != nil {
		fmt.Println(err)
	}
	return err
}
