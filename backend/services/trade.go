package services

import (
	"errors"

	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
)

func ListGoatForSale(farmer *models.Farmer, goatID, price uint) (*models.Transaction, error) {
	var goat models.Goat
	if err := database.DB.First(&goat, goatID).Error; err != nil {
		return nil, err
	}

	if goat.FarmerID != farmer.ID {
		return nil, errors.New("farmer does not own this goat")
	}

	goat.Price = price
	goat.ForSale = true

	if err := database.DB.Save(&goat).Error; err != nil {
		return nil, err
	}
	tx, err := createTx(farmer.ID, goat.ID, price)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func createTx(farmerID, goatID, price uint) (*models.Transaction, error) {
	transaction := models.Transaction{
		SellerID: farmerID,
		GoatID:   goatID,
		Price:    price,
		Status:   "pending",
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func PurchaseGoat(buyerID, transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := database.DB.First(&transaction, transactionID).Error; err != nil {
		return nil, err
	}

	if transaction.Status != "pending" {
		return nil, errors.New("transaction is not pending")
	}

	var buyer, seller models.Farmer
	if err := database.DB.First(&buyer, buyerID).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(&seller, transaction.SellerID).Error; err != nil {
		return nil, err
	}

	if buyer.Wallet < transaction.Price {
		return nil, errors.New("buyer does not have enough funds")
	}

	var goat models.Goat
	if err := database.DB.First(&goat, transaction.GoatID).Error; err != nil {
		return nil, err
	}

	goat.FarmerID = buyerID
	goat.ForSale = false
	buyer.Wallet -= transaction.Price
	seller.Wallet += transaction.Price

	if err := database.DB.Save(&goat).Error; err != nil {
		return nil, err
	}
	if err := database.DB.Save(&buyer).Error; err != nil {
		return nil, err
	}
	if err := database.DB.Save(&seller).Error; err != nil {
		return nil, err
	}

	transaction.Status = "completed"
	transaction.BuyerID = buyerID
	return &transaction, database.DB.Save(&transaction).Error
}

func GetAvailableGoats() ([]models.Goat, error) {
	var goats []models.Goat
	err := database.DB.Where("for_sale = ?", true).Find(&goats).Error
	return goats, err
}
