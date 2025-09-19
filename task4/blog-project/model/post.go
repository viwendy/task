package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint   `json:"user_id"`
	User    User
}

func (p *Post) TableName() string {
	return "posts"
}
