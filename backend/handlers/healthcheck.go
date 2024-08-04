package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/services"
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

func BookAppointment(c *gin.Context) {

	var request struct {
		GoatID uint `json:"goatid" binding:"required"`
		VetID  uint `json:"vetid" binding:"required"`
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

	var goat models.Goat
	if err := database.DB.Where("id = ? AND farmer_id = ?", request.GoatID, farmer.ID).First(&goat).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Goat does not belong to the farmer"})
		return
	}

	appointment := services.CreateAppointment(farmer.ID, request.GoatID, request.VetID)
	if err := database.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func ListAppointment(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	vet, ok := user.(*models.Vet)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a farmer"})
		return
	}

	appointments, err := services.GetAppointments(0, vet.ID, 0, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}
