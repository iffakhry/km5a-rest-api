package routes

import (
	"rest/mvc/controllers"
	"rest/mvc/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] status=${status} method=${method} uri=${uri} latency=${latency_human} \n",
	}))

	e.POST("/login", controllers.LoginController)

	e.GET("/users", controllers.GetUserController, middlewares.JWTMiddleware())
	e.POST("/users", controllers.AddUserController)
	e.POST("/users/alamat", controllers.AddUserAlamatController)

	product := e.Group("/products", middlewares.JWTMiddleware())
	product.GET("", controllers.GetProductController)
	product.POST("", controllers.AddProductController)
	return e
}
