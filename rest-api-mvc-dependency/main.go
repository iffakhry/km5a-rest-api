package main

import (
	"fmt"
	"rest/mvc-di/configs"
	"rest/mvc-di/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var cfg = configs.InitConfig()
	dbMysql := configs.InitMysqlConn(cfg)
	routes.InitRouter(e, dbMysql)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
	// e.Logger.Fatal(e.Start(":8000"))
}
