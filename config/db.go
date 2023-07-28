package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Final/models"
	"Final/utils"
	"fmt"
)

func ConnectDatabase() *gorm.DB {
	username := utils.Getenv("MYSQLDATABASE", "root")
	password := utils.Getenv("MYSQLPASSWORD", "8J2b5pX9RygNKe5iXC0s")
	host := utils.Getenv("MYSQLHOST", "containers-us-west-116.railway.app")
	port := utils.Getenv("MYSQLPORT", "7659")
	database := utils.Getenv("MYSQLDATABASE", "railway")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Category{}, &models.Game{}, &models.Rating{}, &models.Review{}, &models.Comment{}, &models.User{})

	return db
}
