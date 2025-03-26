package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"inventory-management-api/models"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("inventory_mg.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.Admin{}, &models.InventoryItem{}, &models.Restock{})
}
