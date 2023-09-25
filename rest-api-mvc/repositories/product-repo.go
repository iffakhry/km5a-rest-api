package repositories

import (
	"errors"
	"rest/mvc/config"
	"rest/mvc/models"
)

func SelectProduct() ([]models.Product, error) {
	//menggunakan db
	var dataproducts []models.Product
	// select * from users;
	tx := config.DB.Preload("User").Find(&dataproducts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return dataproducts, nil
}

func InsertProduct(data models.Product) error {
	tx := config.DB.Create(&data)
	if tx.Error != nil {
		return errors.New("failed insert data")
	}

	return nil
}
