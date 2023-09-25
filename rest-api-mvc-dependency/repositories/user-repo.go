package repositories

import (
	"rest/mvc-di/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Select() ([]models.User, error) {
	//menggunakan db
	var datausers []models.User
	// select * from users;
	tx := ur.db.Order("created_at desc").Find(&datausers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datausers, nil
}
