package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Final/models"
	"Final/utils"
	"fmt"
)

func ConnectDatabase() *gorm.DB {
	username := utils.Getenv("root", "root")
	password := utils.Getenv("8J2b5pX9RygNKe5iXC0s", "Antimaling2@")
	host := utils.Getenv("containers-us-west-116.railway.app", "127.0.0.1")
	port := utils.Getenv("7659", "3306")
	database := utils.Getenv("railway", "db_game")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Category{}, &models.Game{}, &models.Rating{}, &models.Review{}, &models.Comment{}, &models.User{})

	return db
}
