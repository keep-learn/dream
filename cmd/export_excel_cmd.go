package cmd

import (
	"dream/business/service"
	"github.com/spf13/cobra"
)

// ExportExcelCmd  导出excel的命令
var ExportExcelCmd = &cobra.Command{
	Use:   "export",
	Short: `输入需要导出的内容(eg: "X 信息、Y 信息；甲类、乙类")`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// 初始化依赖扩展
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 启动入口
		service.RunExportExcel(args)
	},
}
