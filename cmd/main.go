package main

import (
	"flag"
	"fmt"
)

// 用户输入
var exportContent string

func init() {
	// 获取用户输入
	flag.StringVar(&exportContent, "export", "", "请输入需要导出的内容！")
}

func main(){
	flag.Parse()

	if exportContent == ""{
		// todo log  异常需要打印日志
	}

	// todo 校验参数合法性
	fmt.Println("inputContent:", exportContent)
}