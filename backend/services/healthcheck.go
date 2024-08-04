package services

import (
	"github.com/vaibhavsijaria/TGC-be.git/database"
	"github.com/vaibhavsijaria/TGC-be.git/models"
)

func CreateAppointment(farmerid, goatid, vetid uint) models.Appointment {
	appointment := models.Appointment{
		FarmerID: farmerid,
		GoatID:   goatid,
		VetID:    vetid,
		Status:   "pending",
	}

	return appointment
}

func GetAppointments(farmerid, vetid, goatid uint, status string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	query := database.DB.Preload("Farmer").Preload("Vet")

	if farmerid != 0 {
		query = query.Where("farmer_id = ?", farmerid)
	}

	if vetid != 0 {
		query = query.Where("vet_id = ?", vetid)
	}

	if goatid != 0 {
		query = query.Where("goat_id = ?", goatid)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}
