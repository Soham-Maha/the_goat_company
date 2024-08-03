package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Farmer struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique"`
	Password   string `json:"password"`
	Wallet     uint
	Goats      []Goat
	Invesments []Invesment
}

type Investor struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique"`
	Password   string `json:"password"`
	Invesments []Invesment
}

type Vet struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	Password     string `json:"password"`
	HealthChecks []HealthCheck
}
