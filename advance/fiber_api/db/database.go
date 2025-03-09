package db

import (
	"fiber_api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

var DatabaseRef Database

func Connect() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Add migrations
	db.AutoMigrate(&models.User{})

	DatabaseRef = Database{Db: db}
}
