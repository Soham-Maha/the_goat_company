package services

import (
	"errors"
	"fmt"

	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
)

func AddMoney(userType string, userID uint, amount uint) error {
	switch userType {
	case "farmer":
		return addMoneyToFarmer(userID, amount)
	default:
		return errors.New("Invalid Type")
	}
}

func addMoneyToFarmer(userID uint, amount uint) error {
	var farmer models.Farmer
	if err := database.DB.First(&farmer, userID).Error; err != nil {
		return fmt.Errorf("error finding farmer: %w", err)
	}
	farmer.Wallet += amount
	if err := database.DB.Save(&farmer).Error; err != nil {
		return fmt.Errorf("error saving investor: %w", err)
	}
	return nil
}

func SubMoney(userType string, userID uint, amount uint) error {
	switch userType {
	case "farmer":
		return subMoneyFromFarmer(userID, amount)
	default:
		return errors.New("Invalid Type")
	}
}

func subMoneyFromFarmer(userID uint, amount uint) error {
	var farmer models.Farmer
	if err := database.DB.First(&farmer, userID).Error; err != nil {
		return fmt.Errorf("error finding farmer: %w", err)
	}
	if farmer.Wallet < amount {
		return errors.New("insufficient funds")
	}
	farmer.Wallet -= amount
	if err := database.DB.Save(&farmer).Error; err != nil {
		return fmt.Errorf("error saving farmer: %w", err)
	}
	return nil
}
