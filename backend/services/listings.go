package services

import (
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
)

func GetGoats(species, sex string, age, price, farmerID uint, forSale *bool) ([]models.Goat, error) {
	var goats []models.Goat
	query := database.DB

	if species != "" {
		query = query.Where("species = ?", species)
	}

	if age != 0 {
		query = query.Where("age = ?", age)
	}

	if sex != "" {
		query = query.Where("sex = ?", sex)
	}

	if price != 0 {
		query = query.Where("price = ?", price)
	}

	if forSale != nil {
		query = query.Where("for_sale = ?", *forSale)
	}

	if farmerID != 0 {
		query = query.Where("farmer_id = ?", farmerID)
	}

	if err := query.Find(&goats).Error; err != nil {
		return nil, err
	}

	return goats, nil
}
