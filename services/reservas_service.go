package services

import (
	"workspace/models"
	"workspace/repositories"
)

func GetReservas() ([]models.Reservas, error) {
	return repositories.GetReservas()
}
