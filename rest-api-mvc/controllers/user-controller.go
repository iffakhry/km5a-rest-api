package controllers

import (
	"net/http"
	"rest/mvc/controllers/responses"
	"rest/mvc/helpers"
	"rest/mvc/middlewares"
	"rest/mvc/models"
	"rest/mvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	// //menggunakan db
	// var datausers []models.User
	// // select * from users;
	// tx := config.DB.Find(&datausers)

	idToken := middlewares.ExtractTokenUserId(c)
	if idToken != 1 {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	responseData, err := repositories.SelectUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error get data"))
	}

	var usersResponse = []responses.UserResponse{}
	for _, value := range responseData {
		usersResponse = append(usersResponse, responses.UserResponse{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success recieve user data", usersResponse))
}

func AddUserController(c echo.Context) error {

	var datausers models.User
	// membaca data yang dikirimkan client/FE
	errBind := c.Bind(&datausers)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind "+errBind.Error()))
	}
	// //menggunakan db
	// // insert into users()values ....
	// tx := config.DB.Create(&datausers)
	err := repositories.InsertUser(datausers)
	if err != nil {
		// response message gagal
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error insert data"))
	}

	//response message berhasil
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success insert data"))
}
