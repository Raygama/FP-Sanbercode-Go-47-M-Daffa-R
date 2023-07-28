package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Final/models"
	"fmt"
)

func ConnectDatabase() *gorm.DB {
	username := "root"
	password := "8J2b5pX9RygNKe5iXC0s"
	host := "containers-us-west-116.railway.app"
	database := "railway"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Category{}, &models.Game{}, &models.Rating{}, &models.Review{}, &models.Comment{}, &models.User{})

	return db
}
