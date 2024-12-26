package config

import (
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	JWTKey = []byte("your-secret-key")
)
