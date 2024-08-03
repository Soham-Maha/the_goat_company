package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
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

	database.DB.Preload("Farmer").First(&invesment, invesment.ID)
	database.DB.Preload("Investor").First(&invesment, invesment.ID)

	c.JSON(http.StatusOK, invesment)
}
