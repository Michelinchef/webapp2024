package model

import "time"

type Post struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	UserID    uint       `json:"user_id"`
	User      User       `json:"user"`
	Comments  []Comment  `json:"comments"`
}
