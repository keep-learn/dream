package service

import (
	"context"
	"fmt"
	"testing"
)

// Construct
func TestUserInput_Construct(t *testing.T) {
	ctx := context.Background()
	input := []string{"X 信息、Y 信息；甲类、乙类、丁类"}
	inputService := NewInput(input)
	err := inputService.Check(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	result := inputService.Construct()

	if fmt.Sprintf("%v", result) != "[{X 信息 0 甲类 1 X 信息-甲类-1} { 0 乙类 2 X 信息-乙类-2} { 0 丁类 3 X 信息-丁类-3} {Y 信息 0 甲类 4 Y 信息-甲类-1} { 0 乙类 5 Y 信息-乙类-2} { 0 丁类 6 Y 信息-丁类-3}]" {
		t.Error("失败")
		return
	}
	t.Log("success")
}
