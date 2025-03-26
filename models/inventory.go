package models

import "time"

// Each inventory item should include fields like name, description, quantity, and timestamps for tracking creation, updates, and last restock.

type InventoryItem struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Quantity       int       `json:"quantity"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	LastRestockAt  time.Time `json:"last_restock_at"`
	RestockHistory []Restock `json:"restock_history"`
}


// For restocking items, ensure the restock amount is validated to be between 10 and 1000.
// Additionally, enforce a rate limit: if an item is restocked more than 3 times within a 24-hour period, the API must return a 429 Too Many Requests error.
// Restock has Amount and it has validation to be between 10-1000

type Restock struct {
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
