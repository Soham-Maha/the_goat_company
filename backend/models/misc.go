package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	SellerID uint
	BuyerID  uint
	GoatID   uint
	Price    uint
	Status   string // pending/completed/cancelled
}
