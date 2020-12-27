package model

import (
	"fmt"
	"lithum/pkg/constvar"
)

type SysField struct {
	ID           uint64 `gorm:"primary;AUTO_INCREMENT;column:id"`
	Label        string `gorm:"column:label"`
	Tag          string `gorm:"column:tag"`
	DefaultValue string `gorm:"column:default_value"`
	Disable      bool   `gorm:"column:disable"`
	OtherOptions string `gorm:"column:options"`
	SysForm      SysForm
	SysFormID    uint64
}

func (field *SysField) Create() error {
	return DB.Self.Create(&field).Error
}

func DeleteField(id uint64) error {
	return DB.Self.Delete(&SysField{ID: id}).Error
}

func (field *SysField) Update() error {
	return DB.Self.Delete(&field).Error
}

func ListField(search string, offset int, limit int) ([]*SysField, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	items := make([]*SysField, 0)
	var count uint64

	where := fmt.Sprintf("label like '%%%s%%'", search)

	if err := DB.Self.Model(&SysField{}).Where(where).Count(&count).Error; err != nil {
		return items, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&items).Error; err != nil {
		return items, count, err
	}

	return items, count, nil
}
