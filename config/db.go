package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Final/models"
	"fmt"
	"os"
)

func ConnectDatabase() *gorm.DB {
	username := os.Getenv("MYSQLUSER")
	password := os.Getenv("MYSQLPASSWORD")
	host := os.Getenv("MYSQLHOST")
	database := os.Getenv("MYSQLDATABASE")

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Category{}, &models.Game{}, &models.Rating{}, &models.Review{}, &models.Comment{}, &models.User{})

	return db
}
