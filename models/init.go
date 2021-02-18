package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.database"), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database!")
	}

	if err := database.AutoMigrate(&Stock{}); err != nil {
		panic(err.Error())
	}

	DB = database
}
