package cmd

import (
	"context"
	"dream/business/service"
	"dream/pkg/log"
	"github.com/spf13/cobra"
	"time"
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
		ctx := context.WithValue(context.Background(), log.TraceID, time.Now().UnixNano())
		service.RunExportExcel(ctx, args)
	},
}
