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
	userType := claims.UserType

	switch userType {
	case "farmer":
		user = new(models.Farmer)
	case "investor":
		user = new(models.Investor)
	case "vet":
		user = new(models.Vet)
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type"})
		c.Abort()
		return
	}

	if err := database.DB.Where("email = ?", claims.Subject).First(user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Set("userType", userType)
	c.Set("claims", claims)
	c.Next()
}
