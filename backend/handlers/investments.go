package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/services"
)

func CreateInvestment(c *gin.Context) {
	invesment := new(models.Invesment)

	if err := c.ShouldBindJSON(invesment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	investor, ok := user.(*models.Investor)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a Investor"})
		return
	}

	invesment.InvestorID = investor.Model.ID

	if err := database.DB.Create(invesment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create investment"})
	}

	if err := services.AddMoney("farmer", invesment.FarmerID, invesment.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update farmer's wallet"})
		return
	}

	database.DB.Preload("Farmer").First(&invesment, invesment.ID)
	database.DB.Preload("Investor").First(&invesment, invesment.ID)

	c.JSON(http.StatusOK, invesment)
}

func ViewAllInvestment(c *gin.Context) {
	var invesments []models.Invesment
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	investor, ok := user.(*models.Investor)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a Investor"})
		return
	}

	if err := database.DB.Where("investor_id = ?", investor.ID).Preload("Farmer").Find(&invesments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve investments"})
		return
	}

	c.JSON(http.StatusOK, invesments)

}
