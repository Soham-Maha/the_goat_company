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

func SeekInvestment(farmerid, amount uint, psplit float32) (*models.Invesment, error) {
	investment := models.Invesment{
		FarmerID:    farmerid,
		Amount:      amount,
		ProfitSplit: psplit,
		Status:      "pending",
	}
	investment.InvestorID = 1
	if err := database.DB.Create(&investment).Error; err != nil {
		return nil, err
	}

	return &investment, nil
}

func AcceptInvestment(investmentid, investorid uint) (*models.Invesment, error) {
	var investment models.Invesment

	if err := database.DB.First(&investment, investmentid).Error; err != nil {
		return nil, err
	}

	if investment.Status != "pending" {
		return nil, errors.New("The investment is already accepted")
	}
	if err := AddMoney("farmer", investment.FarmerID, investment.Amount); err != nil {
		return nil, err
	}
	investment.InvestorID = investorid
	investment.Status = "completed"
	return &investment, database.DB.Save(&investment).Error
}

func GetInvestments(amount uint, profitSplit float32, farmerID, investorID uint, status string) ([]models.Invesment, error) {
	var investments []models.Invesment
	query := database.DB.Preload("Farmer").Preload("Investor")

	if amount != 0 {
		query = query.Where("amount = ?", amount)
	}

	if profitSplit != 0 {
		query = query.Where("profit_split = ?", profitSplit)
	}

	if farmerID != 0 {
		query = query.Where("farmer_id = ?", farmerID)
	}

	if investorID != 0 {
		query = query.Where("investor_id = ?", investorID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&investments).Error; err != nil {
		return nil, err
	}

	return investments, nil
}
