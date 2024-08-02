package models

import "gorm.io/gorm"

type Goat struct {
	gorm.Model
	Species     string `json:"species" binding:"required"`
	Age         uint   `json:"age" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Description string `json:"description"`
	FarmerID    uint
	Farmer      Farmer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
