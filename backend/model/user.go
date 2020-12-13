package model

import (
	"fmt"

	"lithum/pkg/auth"
	"lithum/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
)

// User represents a registered user.
type User struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Roles []Role `gorm:"many2many:users_roles;'"`
}

// Create creates a new user account.
func (user *User) Create() error {
	return DB.Self.Create(&user).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := User{}
	user.ID = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (user *User) Update() error {
	return DB.Self.Save(user).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*User, error) {
	user := &User{}
	d := DB.Self.Where("username = ?", username).First(&user)
	return user, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*User, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*User, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (user *User) Compare(pwd string) (err error) {
	err = auth.Compare(user.Password, pwd)
	return
}

// Encrypt the user password.
func (user *User) Encrypt() (err error) {
	user.Password, err = auth.Encrypt(user.Password)
	return
}

// Validate the fields.
func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)
}
