package routes

import (
	"rest/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.AddUserController)

	product := e.Group("/products")
	product.GET("", controllers.GetProductController)
	product.POST("", controllers.GetUserController)
	return e
}
