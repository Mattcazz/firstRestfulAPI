package repositories

import (
	"errors"
	"workspace/db"
	"workspace/models"
)

func GetAlojamientos() ([]models.Alojamientos, error) {

	var alojamientos []models.Alojamientos

	err := db.DB.Preload("Reservas").Preload("Reservas.Formalizacion").Find(&alojamientos).Error

	return alojamientos, err
}

func GetAlojamientoById(id uint) (models.Alojamientos, error) {
	var alojamiento models.Alojamientos

	res := db.DB.Where("idalojamiento = ?", id).Preload("Reservas").Find(&alojamiento)

	if res.RowsAffected == 0 {
		return alojamiento, errors.New("id does not exist in the Data Base")
	}

	return alojamiento, res.Error
}

func CreateAlojamiento(alojamiento *models.Alojamientos) error {
	return db.DB.Create(alojamiento).Error
}
