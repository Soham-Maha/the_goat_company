package models

import "gorm.io/gorm"

type Goat struct {
	gorm.Model
	Species     string `json:"species" binding:"required"`
	Age         uint   `json:"age" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Description string `json:"description"`
	ImageURL    string
	FarmerID    uint
	Farmer      Farmer
}

type Invesment struct {
	gorm.Model
	Amount      uint    `json:"amount" binding:"required"`
	ProfitSplit float32 `json:"psplit" binding:"required"`
	FarmerID    uint    `json:"farmerid" binding:"required"`
	Farmer      Farmer
	InvestorID  uint
	Investor    Investor
}
