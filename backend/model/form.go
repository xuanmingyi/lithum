package model

import (
	"fmt"
	"lithum/pkg/constvar"
)

type SysForm struct {
	ID            uint64 `gorm:"primary;AUTO_INCREMENT;column:id"`
	Title         string `gorm:"column:title"`
	Size          string `gorm:"column:size"`
	LabelPosition string `gorm:"column:label_position"`
	LabelWidth    uint8  `gorm:"column:label_width"`
	Span          uint8  `gorm:"span"`
	Gutter        uint8  `gorm:"gutter"`
	FormBtns      bool   `gorm:"form_btns"`
	Disabled      bool   `gorm:"disabled"`
	Fields        []SysField
}

func (form *SysForm) Create() error {
	return DB.Self.Create(&form).Error
}

func DeleteForm(id uint64) error {
	return DB.Self.Delete(&SysForm{ID: id}).Error
}

func (form *SysForm) Update() error {
	return DB.Self.Update(&form).Error
}

func ListForm(search string, offset int, limit int) ([]*SysForm, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	items := make([]*SysForm, 0)
	var count uint64

	where := fmt.Sprintf("title like '%%%s%%'", search)
	if err := DB.Self.Model(&SysForm{}).Where(where).Count(&count).Error; err != nil {
		return items, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&items).Error; err != nil {
		return items, count, err
	}

	return items, count, nil
}
