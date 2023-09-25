package controllers

import (
	"log"
	"net/http"
	"rest/mvc/controllers/requests"
	"rest/mvc/helpers"
	"rest/mvc/middlewares"
	"rest/mvc/models"
	"rest/mvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetProductController(c echo.Context) error {
	qLimit := c.QueryParam("limit")
	qPage := c.QueryParam("page")
	log.Println("limit:", qLimit, "page:", qPage)

	response, err := repositories.SelectProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success recieve product data",
		"data":    response,
	})
}

func AddProductController(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(requests.ProductRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind "+errBind.Error()))
	}
	// log.Println("product", input)
	// mapping dari struct request ke stuct model gorm
	data := models.Product{
		Name:   input.Name,
		Stock:  uint(input.Stock),
		UserID: uint(idToken),
	}
	err := repositories.InsertProduct(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error insert data "+err.Error()))
	}

	//response message berhasil
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success insert data"))
}
