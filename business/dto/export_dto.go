package dto

// 导出的结构体
type ExportItem struct {
	Rank     string // 等级
	Seq      int    // 序号
	Type     string // 类型
	Priority int    // 优先级
	Desc     string // 说明
}

// 导出结果信息
type ExportResult struct {
	ExcelPath string // 导出excel的path信息
}
