package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/utils"
	"golang.org/x/crypto/bcrypt"
)

type request struct {
	UserType string `json:"usertype"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name not provided"})
		return
	}

	var user interface{}
	var existingUser interface{}

	switch req.UserType {
	case "farmer":
		user = &models.Farmer{
			Email: req.Email,
			Name:  req.Name,
		}
		existingUser = &models.Farmer{}
	case "investor":
		user = &models.Investor{
			Email: req.Email,
			Name:  req.Name,
		}
		existingUser = &models.Investor{}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
		return
	}

	if err := database.DB.Where("email = ?", req.Email).First(existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	switch u := user.(type) {
	case *models.Farmer:
		u.Password = string(hash)
	case *models.Investor:
		u.Password = string(hash)
	}

	if err := database.DB.Create(user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user interface{}

	switch req.UserType {
	case "farmer":
		user = &models.Farmer{}
	case "investor":
		user = &models.Investor{}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
		return
	}

	if err := database.DB.Where("email = ?", req.Email).First(user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var hashedPassword string
	switch u := user.(type) {
	case *models.Farmer:
		hashedPassword = u.Password
	case *models.Investor:
		hashedPassword = u.Password
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func LoginCheck(c *gin.Context) {

	user, _ := c.Get("user")
	userType, _ := c.Get("userType")
	if userType == "farmer" {
		farmer, _ := user.(*models.Farmer)
		database.DB.Preload("Goats").First(farmer, farmer.ID)
		database.DB.Preload("Invesments").First(farmer, farmer.ID)
		c.JSON(200, gin.H{
			"message": "You are logged in",
			"user":    farmer,
		})
	} else if userType == "investor" {
		investor, _ := user.(*models.Investor)
		database.DB.Preload("Invesments").First(investor, investor.ID)
		c.JSON(200, gin.H{
			"message": "You are logged in",
			"user":    investor,
		})
	}

}
