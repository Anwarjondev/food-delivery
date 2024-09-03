package utils

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
	"food-delivery/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("fooddelivery.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    
    if err := DB.AutoMigrate(&models.User{}, &models.Courier{}, &models.Order{}); err != nil {
        log.Fatal("Failed to auto-migrate database:", err)
    }
}
