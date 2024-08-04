package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/services"
)

func CreateGoat(c *gin.Context) {
	goat := new(models.Goat)

	// if err := c.ShouldBindJSON(goat); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	goat.Species = c.Request.FormValue("species")
	goat.Description = c.Request.FormValue("description")

	goatAgeStr := c.Request.FormValue("age")
	goatAge, err := strconv.ParseUint(goatAgeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age value"})
		return
	}
	goat.Age = uint(goatAge)

	priceStr := c.Request.FormValue("price")
	price, err := strconv.ParseUint(priceStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age value"})
		return
	}
	goat.Price = uint(price)

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
	goat.FarmerID = farmer.Model.ID

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	filePath := fmt.Sprintf("./uploads/goats/%s_%s", uuid.New().String(), file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	goat.ImageURL = filePath

	if err := database.DB.Create(goat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list goat"})
		return
	}
	database.DB.Preload("Farmer").First(&goat, goat.ID)
	c.JSON(http.StatusOK, goat)
}

func ListGoats(c *gin.Context) {
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

	if sex := c.Query("sex"); sex != "" {
		query = query.Where("sex = ?", sex)
	}

	if age := c.Query("age"); age != "" {
		query = query.Where("age = ?", age)
	}

	query = query.Where("for_sale = ?", true)

	if err := query.Find(&Goats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Goats)
}

func GetMyListings(c *gin.Context) {
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

	forsale := true
	goats, err := services.GetGoats("", "", 0, 0, farmer.ID, &forsale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, goats)
}
