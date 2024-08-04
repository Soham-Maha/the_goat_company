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

type Appointment struct {
	gorm.Model
	FarmerID      uint
	Farmer        Farmer
	GoatID        uint
	Goat          Goat
	VetID         uint
	Vet           Vet
	Status        string
	HealthCheckID *uint
	HealthCheck   HealthCheck
}
