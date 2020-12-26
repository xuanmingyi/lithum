package service

import (
	"lithum/model"
	. "lithum/model/response/form"
)

func ListForm(name string, offset, limit int) (items []*FormInfo, count uint64, err error) {
	var forms []*model.SysForm
	forms, count, err = model.ListForm(name, offset, limit)
	for _, item := range forms {
		items = append(items, &FormInfo{
			ID:            item.ID,
			Title:         item.Title,
			Size:          item.Size,
			LabelPosition: item.LabelPosition,
			LabelWidth:    item.LabelWidth,
			Span:          item.Span,
			Gutter:        item.Gutter,
			FormBtns:      item.FormBtns,
			Disabled:      item.Disabled,
		})
	}
	return
}
