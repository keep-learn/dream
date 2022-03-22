package service

import (
	"errors"
	"strings"
)

// 处理用户输入的信息
type UserInput struct {
	Input   string   // 输入的字符串
	Ranking []string // 等级
	Types   []string // 类型
}

func NewInput(input string) UserInput {
	return UserInput{Input: input}
}

// check 用户的输入
// todo 安全相关的调试
func (ui UserInput) Check() (err error) {
	// todo 暂不考虑并发问题
	if ui.Input == "" {
		// todo warn 日志
		return errors.New("输入内容为空")
	}

	inputArr := strings.Split(ui.Input, "；")
	if len(inputArr) != 2 {
		return errors.New("需要输入等级和类型")
	}
	// todo 数据去重
	// todo 合法数据的校验

	ui.Ranking = strings.Split(inputArr[0], "、")
	ui.Types = strings.Split(inputArr[1], "、")
	return
}
