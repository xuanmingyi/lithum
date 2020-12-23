package model

type SysField struct {
	ID uint64
	Label string
	Tag string
	DefaultValue string
	Disable bool
	OtherOptions string
}

type SysForm struct {
	ID uint64 `gorm:"primary;AUTO_INCREMENT;column:id"`
	Size string `gorm:"column:size"`
	LabelPosition string `gorm:"column:label_position"`
	LabelWidth string `gorm:"column:label_width"`
	Span string	`gorm:"span"`
	Gutter string `gorm:"gutter"`
	FormBtns bool `gorm:"form_btns"`
	Disabled bool `gorm:"disabled"`
	Fields []SysField
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