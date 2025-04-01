package services

import (
	"errors"
	"workspace/models"
	"workspace/repositories"
)

func GetAlojamientos() ([]models.Alojamientos, error) {
	return repositories.GetAlojamientos()
}

func GetALojamientoById(id uint) (models.Alojamientos, error) {
	return repositories.GetAlojamientoById(id)
}

func CreateAlojamiento(alojamiento *models.Alojamientos) error {
	if alojamiento.Propietario == "" {
		return errors.New("Propietario arg is missing")
	}

	if alojamiento.Ciudad == "" {
		return errors.New("Ciudad arg is missing")
	}

	return repositories.CreateAlojamiento(alojamiento)

}
