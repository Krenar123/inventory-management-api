package controllers

import (
	"inventory-management-api/db"
	"inventory-management-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var item models.InventoryItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidRequest})
		return
	}
	db.DB.Create(&item)
	c.JSON(http.StatusOK, item)
}

func ListItems(c *gin.Context) {
	lowStock := c.Query("low_stock")
	var items []models.InventoryItem

	if lowStock == "true" {
		db.DB.Where("quantity <= ?", 20).Find(&items)
	} else {
		db.DB.Find(&items)
	}
	c.JSON(http.StatusOK, items)
}

func RestockItem(c *gin.Context) {
	var request struct {
		Amount int `json:"amount"`
	}
	id := c.Param("id")
	if err := c.ShouldBindJSON(&request); err != nil || request.Amount < 10 || request.Amount > 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrAmountRange})
		return
	}

	var item models.InventoryItem
	if err := db.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": ErrItemNotFound})
		return
	}

	// checking restock rate limit (should be max 3 per 24 hours)
	var count int64
	since := time.Now().Add(-24 * time.Hour)
	db.DB.Model(&models.Restock{}).
		Where("inventory_item_id = ? AND created_at >= ?", item.ID, since).
		Count(&count)

	if count >= 3 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": ErrRestockLimitExceeded})
		return
	}

	// updating quantity on item and logging the restock
	item.Quantity += request.Amount
	db.DB.Save(&item)

	restock := models.Restock{
		InventoryItemID: item.ID,
		Amount:          request.Amount,
	}
	db.DB.Create(&restock)

	c.JSON(http.StatusOK, item)
}

func RestockHistory(c *gin.Context) {
	id := c.Param("id")
	var restocks []models.Restock
	if err := db.DB.Where("inventory_item_id = ?", id).
		Order("created_at desc").
		Find(&restocks).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": ErrItemNotFound})
		return
	}
	c.JSON(http.StatusOK, restocks)
}
