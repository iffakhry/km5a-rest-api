package config

import (
	"fmt"
	"rest/mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "qwerty123",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "db_km5_gorm",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}

func InitDBTest() {
	const DB_USER_TEST = "root"
	const DB_PASS_TEST = "qwerty123"
	const DB_HOST_TEST = "127.0.0.1"
	const DB_PORT_TEST = "3306"
	const DB_NAME_TEST = "db_km5_gorm_test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
	DB.Migrator().DropTable(&models.Product{})
	DB.AutoMigrate(&models.Product{})
}
