package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/utils"
)

func AuthMiddleware(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	var user interface{}
	var userType string
	if err := database.DB.Where("email = ?", claims.Subject).First(&models.Farmer{}).Error; err == nil {
		user = new(models.Farmer)
		userType = "farmer"
	} else if err := database.DB.Where("email = ?", claims.Subject).First(&models.Investor{}).Error; err == nil {
		user = new(models.Investor)
		userType = "investor"
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}
	database.DB.Where("email = ?", claims.Subject).First(user)

	c.Set("user", user)
	c.Set("userType", userType)
	c.Set("claims", claims)
	c.Next()
}
