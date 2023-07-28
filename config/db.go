package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Final/models"
	"fmt"
)

func ConnectDatabase() *gorm.DB {
	username := "root"
	password := "Antimaling2@"
	host := "tcp(0.0.0.0:3306)"
	database := "db_game"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Category{}, &models.Game{}, &models.Rating{}, &models.Review{}, &models.Comment{}, &models.User{})

	return db
}
