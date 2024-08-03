package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
	"github.com/vaibhavsijaria/TGC-be.git/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}
}

func main() {
	database.Init()

	if err := database.RunMigrations(&models.Farmer{}, &models.Investor{}, &models.Goat{}, &models.Invesment{}); err != nil {
		os.Exit(1)
	}

	r := gin.Default()

	routes.UserRoutes(r)
	routes.ListingRoutes(r)
	routes.InvesmentRoute(r)
	r.Run()
}
