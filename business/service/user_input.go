package service

import (
	"context"
	"dream/business/dto"
	"dream/pkg/log"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

// UserInput 处理用户输入的信息
type UserInput struct {
	Input   []string // 输入的字符串
	Ranking []string // 等级
	Types   []string // 类型
}

func NewInput(args []string) *UserInput {
	return &UserInput{Input: args}
}

// Check check 用户的输入
// todo 安全相关的调试
func (ui *UserInput) Check(ctx context.Context) (err error) {
	log.New().WithContext(ctx).Info("start check", zap.Any("input", ui.Input))
	// todo 暂不考虑并发问题
	if len(ui.Input) != 1 {
		log.New().WithContext(ctx).Warn("请输入正确的参数")
		return errors.New("请输入正确的参数")
	}
	inputArr := strings.Split(ui.Input[0], "；")
	if len(inputArr) != 2 {
		log.New().WithContext(ctx).Warn("需要输入等级和类型")
		return errors.New("需要输入等级和类型")
	}

	// todo 数据合法校验细节
	ui.Ranking = strings.Split(inputArr[0], "、")
	ui.Types = strings.Split(inputArr[1], "、")
	if len(ui.Ranking) == 0 || len(ui.Types) == 0 {
		log.New().WithContext(ctx).Warn("等级或类型信息异常")
		return errors.New("等级或类型信息异常，请检查（请使用中文符合'、' 和 '；'）")
	}
	return
}

// Construct 构造输出的数据
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
			tmp.Desc = fmt.Sprintf("%s-%s操作-%d", rank, typeInfo, typeIndex)
			seq++
			//typeIndex++
			data = append(data, tmp)
		}
	}

	return
}
