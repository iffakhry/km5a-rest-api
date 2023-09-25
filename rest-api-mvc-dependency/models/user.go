package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        string `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
