package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
)

type checkup struct {
	Status string `json:"status" binding:"required"`
	Notes  string `json:"notes" binding:"required"`
	GoatID uint   `json:"goatid" binding:"required"`
}

func HealthCheck(c *gin.Context) {
	var check checkup
	healthCheck := new(models.HealthCheck)

	if err := c.ShouldBindJSON(&check); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	healthCheck.GoatID = check.GoatID
	healthCheck.Notes = check.Notes
	healthCheck.Status = check.Status

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	vet, ok := user.(*models.Vet)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a vet"})
		return
	}

	var goat models.Goat
	if err := database.DB.First(&goat, healthCheck.GoatID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Goat not found"})
		return
	}

	healthCheck.VetID = vet.ID
	if err := database.DB.Create(&healthCheck).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create health check"})
		return
	}

	database.DB.Preload("Goat").First(&healthCheck, healthCheck.ID)
	database.DB.Preload("Vet").First(&healthCheck, healthCheck.ID)
	database.DB.Preload("Farmer").First(&healthCheck.Goat, healthCheck.Goat.ID)

	c.JSON(http.StatusOK, healthCheck)
}
