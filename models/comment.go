package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint
	PostID  uint
	User    User
	Post    Post
}
