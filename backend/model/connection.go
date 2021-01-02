package model

import (
	"fmt"
	"lithum/pkg/constvar"
	"time"
)

type Connection struct {
	ID        uint64    `gorm:"primary;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:create_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:update_at" json:"updated_at"`
	Name      string    `gorm:"name" json:"name"`
	Display   string    `gorm:"display" json:"display"`
	Config    string    `gorm:"config" json:"config"`
}

func (c *Connection) Create() error {
	return DB.Self.Create(&c).Error
}

func DeleteConnection(id uint64) error {
	return DB.Self.Delete(&Connection{ID: id}).Error
}

func (c *Connection) Update() error {
	return DB.Self.Update(&c).Error
}

func ListConnection(search string, offset int, limit int) ([]*Connection, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	items := make([]*Connection, 0)
	var count uint64

	where := fmt.Sprintf("")

	if err := DB.Self.Model(&Connection{}).Where(where).Count(&count).Error; err != nil {
		return items, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&items).Error; err != nil {
		return items, count, err
	}
	return items, count, nil
}
