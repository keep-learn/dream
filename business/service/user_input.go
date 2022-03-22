package service

import (
	"dream/business/dto"
	"errors"
	"fmt"
	"strings"
)

// 处理用户输入的信息
type UserInput struct {
	Input   string   // 输入的字符串
	Ranking []string // 等级
	Types   []string // 类型
}

func NewInput(input string) *UserInput {
	return &UserInput{Input: input}
}

// check 用户的输入
// todo 安全相关的调试
func (ui *UserInput) Check() (err error) {
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

// 构造输出的数据
func (ui *UserInput) Construct() (data []dto.ExportItem) {

	mapTypeWithPriority := map[string]int{
		"甲类": 99,
		"乙类": 88,
		"丁类": 77,
	}

	data = make([]dto.ExportItem, 0)
	// 序号
	seq := 1
	for _, rank := range ui.Ranking {
		// typeIndex 记录类型的次序
		typeIndex := 1
		for index, typeInfo := range ui.Types {
			var tmp dto.ExportItem
			if index == 0 {
				tmp.Rank = rank
			}
			tmp.Seq = seq
			tmp.Type = typeInfo
			if pri, ok := mapTypeWithPriority[typeInfo]; ok {
				tmp.Priority = pri
			}
			tmp.Desc = fmt.Sprintf("%s-%s-%d", rank, typeInfo, typeIndex)
			seq++
			//typeIndex++
			data = append(data, tmp)
		}
	}

	return
}
