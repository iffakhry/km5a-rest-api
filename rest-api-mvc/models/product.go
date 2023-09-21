package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string `json:"name"`
	Stock  uint   `json:"stock"`
	UserID uint   `json:"user_id"`
	User   User
}
