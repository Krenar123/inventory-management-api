package models

import "gorm.io/gorm"

// Each inventory item should include fields like name, description, quantity, and timestamps for tracking creation, updates, and last restock.

type InventoryItem struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Restocks    []Restock
}

type Restock struct {
	gorm.Model
	InventoryItemID uint
	Amount          int
}
