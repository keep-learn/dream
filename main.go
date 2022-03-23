package main

import (
	"dream/cmd"
	"flag"
)

// 用户输入
var exportContent string

func init() {
	// 获取用户输入
	// eg: "X 信息、Y 信息；甲类、乙类、丁类"
	flag.StringVar(&exportContent, "export", "", "请输入需要导出的内容！")
}

func main() {
	cmd.Execute()
}
