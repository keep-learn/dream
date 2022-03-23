package service

import (
	"context"
	"dream/business/dto"
	"testing"
)

func TestNewExportExcel(t *testing.T) {
	contents := []dto.ExportItem{
		{
			Rank:     "X信息",
			Seq:      1,
			Type:     "甲类",
			Priority: 99,
			Desc:     "X 信息-甲类操作-1",
		},
		{
			Rank:     "",
			Seq:      2,
			Type:     "甲类",
			Priority: 99,
			Desc:     "X 信息-甲类操作-2",
		},
	}
	_, err := NewExportExcel().Export(context.Background(), contents)
	if err != nil {
		t.Error(err)
	}
}
