package model

type Role struct {
	BaseModel
	Rolename string `json:"rolename" gorm:"column:rolename;not null"`
	Users[] User `gorm:"many2many:users_roles;'"`
}

func (role *Role) Create() error {
	return DB.Self.Create(&role).Error
}