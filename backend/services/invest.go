package services

import (
	"errors"

	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
)

func InitiateInvestment(farmerid, investorid, amount uint, psplit float32) (*models.Invesment, error) {
	investment := models.Invesment{
		FarmerID:    farmerid,
		InvestorID:  investorid,
		Amount:      amount,
		ProfitSplit: psplit,
		Status:      "pending",
	}

	if err := database.DB.Create(&investment).Error; err != nil {
		return nil, err
	}

	return &investment, nil
}

func RecieveInvestment(farmerid, investmentid uint, status string) (*models.Invesment, error) {
	var investment models.Invesment

	if err := database.DB.First(&investment, investmentid).Error; err != nil {
		return nil, err
	}

	if investment.FarmerID != farmerid {
		return nil, errors.New("Unauthorized")
	}
	switch status {
	case "completed":
		if err := AddMoney("farmer", investment.FarmerID, investment.Amount); err != nil {
			return nil, err
		}
	case "cancelled":
	default:
		return nil, errors.New("Invalid status")
	}
	investment.Status = status
	return &investment, database.DB.Save(&investment).Error
}
