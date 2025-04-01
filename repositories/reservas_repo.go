package repositories

import (
	"workspace/db"
	"workspace/models"

	"gorm.io/gorm/clause"
)

func GetReservas() ([]models.Reservas, error) {
	var reservas []models.Reservas

	err := db.DB.Preload(clause.Associations).Find(&reservas).Error

	return reservas, err
}
