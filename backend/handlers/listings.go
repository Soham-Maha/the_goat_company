package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
)

func CreateGoat(c *gin.Context) {
	goat := new(models.Goat)

	if err := c.ShouldBindJSON(goat); err != nil {
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
	// goat.Farmer = *farmer
	goat.FarmerID = farmer.Model.ID

	if err := database.DB.Model(farmer).Association("Goats").Append(goat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	database.DB.Preload("Goats").First(farmer, farmer.ID)
	c.JSON(http.StatusOK, goat)
}

func ListProducts(c *gin.Context) {
	var Goats []models.Goat
	query := database.DB

	if species := c.Query("species"); species != "" {
		query = query.Where("species = ?", species)
	}

	if minPrice := c.Query("min_price"); minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}

	if maxPrice := c.Query("max_price"); maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	if err := query.Find(&Goats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Goats)
}
