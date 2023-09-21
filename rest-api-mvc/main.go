package main

import (
	"rest/mvc/config"
	"rest/mvc/routes"
)

func main() {
	// membuka connection ke mysql
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))

}
