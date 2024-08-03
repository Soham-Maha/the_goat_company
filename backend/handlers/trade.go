package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/services"
)

func ListGoatForSale(c *gin.Context) {

	var request struct {
		GoatID uint `json:"goatid" binding:"required"`
		Price  uint `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	farmer, ok := user.(*models.Farmer)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a farmer"})
		return
	}
	tx, err := services.ListGoatForSale(farmer, request.GoatID, request.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func PurchaseGoat(c *gin.Context) {
	var request struct {
		TransactionID uint `json:"transactionid" binding:"required"`
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	farmer, ok := user.(*models.Farmer)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a farmer"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := services.PurchaseGoat(farmer.ID, request.TransactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}
