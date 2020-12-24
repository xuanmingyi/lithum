package model

type SysField struct {
	ID           uint64 `gorm:"primary;AUTO_INCREMENT;column:id"`
	Label        string `gorm:"label"`
	Tag          string `gorm:"tag"`
	DefaultValue string `gorm:"default_value"`
	Disable      bool   `gorm:"disable"`
	OtherOptions string
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
