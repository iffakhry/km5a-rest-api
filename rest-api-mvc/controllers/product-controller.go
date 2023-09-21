package controllers

import (
	"net/http"
	"rest/mvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetProductController(c echo.Context) error {
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
