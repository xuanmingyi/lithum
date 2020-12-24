package model

type SysForm struct {
	ID            uint64 `gorm:"primary;AUTO_INCREMENT;column:id"`
	Size          string `gorm:"column:size"`
	LabelPosition string `gorm:"column:label_position"`
	LabelWidth    uint8  `gorm:"column:label_width"`
	Span          uint8  `gorm:"span"`
	Gutter        string `gorm:"gutter"`
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
