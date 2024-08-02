package database

import (
	"log"

	"github.com/vaibhavsijaria/TGC-be.git/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := utils.GetEnv("DB_URL")
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}

func RunMigrations(dst ...interface{}) error {
	log.Println("Starting database migrations...")

	err := DB.AutoMigrate(dst...)
	if err != nil {
		log.Printf("Error during migration: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully.")
	return nil
}
