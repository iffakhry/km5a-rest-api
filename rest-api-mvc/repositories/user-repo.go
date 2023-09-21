package repositories

import (
	"errors"
	"rest/mvc/config"
	"rest/mvc/models"
)

func SelectUsers() ([]models.User, error) {
	//menggunakan db
	var datausers []models.User
	// select * from users;
	tx := config.DB.Find(&datausers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datausers, nil
}

func InsertUser(data models.User) error {
	//menggunakan db
	// insert into users()values ....
	tx := config.DB.Create(&data)
	if tx.Error != nil {
		return errors.New("failed insert data")
	}

	return nil

}
