package models

import (
	"gorm.io/gorm"
)

type Goat struct {
	gorm.Model
	Species      string `json:"species" binding:"required"`
	Age          uint   `json:"age" binding:"required"`
	Sex          string `json:"sex" binding:"required"`
	Price        uint   `json:"price" binding:"required"`
	Description  string `json:"description"`
	ImageURL     string
	ForSale      bool
	FarmerID     uint
	Farmer       Farmer
	HealthChecks []HealthCheck
}

type Invesment struct {
	gorm.Model
	Amount      uint    `json:"amount" binding:"required"`
	ProfitSplit float32 `json:"psplit" binding:"required"`
	FarmerID    uint
	Farmer      Farmer
	InvestorID  uint
	Investor    Investor
	Status      string // "completed" or "pending" or "cancelled"
}

type HealthCheck struct {
	gorm.Model
	GoatID uint `json:"goatid" binding:"required"`
	Goat   Goat
	VetID  uint
	Vet    Vet
	Status string `json:"status" binding:"required"`
	Notes  string `json:"notes" binding:"required"`
}
