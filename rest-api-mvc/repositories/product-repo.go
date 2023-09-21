package repositories

import (
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
