package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"unique;not null" json:"username"`
	Password string  `gorm:"not null" json:"password"`
	Email    *string `json:"email"`
}

func (u *User) TableName() string {
	return "users"
}
