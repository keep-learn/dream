package service

import (
	"context"
	"dream/business/dto"
	"dream/pkg/log"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
	"unicode/utf8"
)

// 等级信息
var mapRanking = map[string]bool{
	"X 信息": true,
	"Y 信息": true,
}

// 类型
var mapTypeWithPriority = map[string]int{
	"甲类": 99,
	"乙类": 88,
	"丁类": 77,
}

// UserInput 处理用户输入的信息
type UserInput struct {
	Input   []string // 输入的字符串
	Ranking []string // 等级
	Types   []string // 类型
}

func NewInput(args []string) *UserInput {
	return &UserInput{Input: args}
}

// Check 用户的输入
func (ui *UserInput) Check(ctx context.Context) (err error) {

	log.New().WithContext(ctx).Info("start check", zap.Any("input", ui.Input))

	// 校验用户输入内容
	inputArr, err := ui.checkInput(ctx)
	if err != nil {
		return
	}
	// 校验等级
	if err = ui.checkRank(ctx, inputArr[0]); err != nil {
		return
	}
	// 校验等级
	if err = ui.checkType(ctx, inputArr[1]); err != nil {
		return
	}
	// 判断参数是否为空
	if len(ui.Ranking) == 0 || len(ui.Types) == 0 {
		return errors.New(`等级或类型信息为空，请检查 (eg: "X 信息、Y 信息；甲类")`)
	}
	return
}

// checkInput 校验输入字段
func (ui *UserInput) checkInput(ctx context.Context) (inputArr []string, err error) {

	if len(ui.Input) != 1 {
		err = errors.New(`请输入正确的参数 (eg: "X 信息、Y 信息；甲类")`)
		return
	}

	// 用户输入最长限制
	maxStrLen := viper.GetInt("safe.maxInputLen")
	if utf8.RuneCountInString(ui.Input[0]) > maxStrLen {
		err = errors.New(fmt.Sprintf("用户输入字符超限，最多%d个字符", maxStrLen))
		return
	}

	inputArr = strings.Split(ui.Input[0], "；")
	if len(inputArr) != 2 {
		err = errors.New(`需要输入等级和类型 (eg: "X 信息、Y 信息；甲类")`)
		return
	}
	return
}

// checkRank 校验等级字段
func (ui *UserInput) checkRank(ctx context.Context, rankStr string) (err error) {
	for _, item := range strings.Split(rankStr, "、") {
		if item != "" {
			if _, ok := mapRanking[item]; !ok {
				msg := fmt.Sprintf("等级信息[%s]不存在，请检查", item)
				err = errors.New(msg)
				return
			}
			ui.Ranking = append(ui.Ranking, item)
		}
	}
	return
}

// checkType 校验类型
func (ui *UserInput) checkType(ctx context.Context, typeStr string) (err error) {
	for _, item := range strings.Split(typeStr, "、") {
		if item != "" {
			if _, ok := mapTypeWithPriority[item]; !ok {
				msg := fmt.Sprintf("类型信息[%s]不存在，请检查", item)
				err = errors.New(msg)
				return
			}
			ui.Types = append(ui.Types, item)
		}
	}
	return
}

// Construct 构造输出的数据
func (ui *UserInput) Construct() (data []dto.ExportItem) {

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
