package cmd

import (
	"dream/pkg/conf"
	"dream/pkg/log"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "dream",
	Short: "一个导出Excel的工具",
	Long: `一个动态读取用户数据并导出EXCEL的命令行工具:
`,
	TraverseChildren: true,
	// 移除 completion 命令
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func init() {

	// 配置文件文件
	RootCmd.PersistentFlags().StringVar(&conf.CfgFile, "c", "config/default.toml", "config file")

	// 提前初始化资源（配置、日志等）
	cobra.OnInitialize(onInitialize)

	// 添加命令的cmd
	RootCmd.AddCommand(ExportExcelCmd)
}

// onInitialize 资源初始化
func onInitialize() {
	conf.Init()
	log.Init()
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		color.Red("异常：%s", err.Error())
		os.Exit(1)
	}
}
