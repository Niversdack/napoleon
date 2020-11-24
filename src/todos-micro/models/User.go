package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}
type UserWithoutPassword struct {
	ID       uint
	Username string
}

func (u *User) ToWithoutPassword() UserWithoutPassword {
	user := UserWithoutPassword{
		ID:       u.Model.ID,
		Username: u.Username,
	}
	return user
}
