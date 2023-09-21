package controllers

import (
	"net/http"
	"rest/mvc/models"
	"rest/mvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {

	// //menggunakan db
	// var datausers []models.User
	// // select * from users;
	// tx := config.DB.Find(&datausers)

	response, err := repositories.SelectUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success recieve user data",
		"data":    response,
	})
}

func AddUserController(c echo.Context) error {

	var datausers models.User
	// membaca data yang dikirimkan client/FE
	errBind := c.Bind(&datausers)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error bind " + errBind.Error(),
		})
	}
	// //menggunakan db
	// // insert into users()values ....
	// tx := config.DB.Create(&datausers)
	err := repositories.InsertUser(datausers)
	if err != nil {
		// response message gagal
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error insert " + err.Error(),
		})
	}

	//response message berhasil
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success insert user data",
	})
}
