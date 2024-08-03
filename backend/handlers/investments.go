package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/services"
)

func OfferInvestment(c *gin.Context) {
	var request struct {
		Farmerid uint    `json:"farmerid" binding:"required"`
		Amount   uint    `json:"amount" binding:"required"`
		Psplit   float32 `json:"psplit" binding:"required"`
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

	investor, ok := user.(*models.Investor)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a Investor"})
		return
	}

	investment, err := services.InitiateInvestment(request.Farmerid, investor.ID, request.Amount, request.Psplit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, investment)

}

func AcceptInvestment(c *gin.Context) {
	var request struct {
		InvestmentId uint   `json:"investmentid" binding:"required"`
		Status       string `json:"status" binding:"required"`
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

	investment, err := services.RecieveInvestment(farmer.ID, request.InvestmentId, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, investment)
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
