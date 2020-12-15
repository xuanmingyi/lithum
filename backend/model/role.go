package model

type Role struct {
	BaseModel
	Rolename string `json:"rolename" gorm:"column:rolename;not null"`
	Users    []User `gorm:"many2many:users_roles;'"`
}

func (role *Role) Create() error {
	return DB.Self.Create(&role).Error
}

func DeleteRole(id uint64) error {
	role := Role{}
	role.ID = id
	return DB.Self.Delete(&role).Error
}

func (role *Role) Update() error {
	return DB.Self.Save(role).Error
}

func GetRole(rolename string) (*Role, error) {
	role := &Role{}
	db := DB.Self.Where("rolename = ?", rolename).First(&role)
	return role, db.Error
}

func ListRole(rolename string, offset, limit int) ([]*Role, uint64, error) {
	return nil, 0, nil
}
