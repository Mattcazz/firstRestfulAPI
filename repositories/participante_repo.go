package repositories

import (
	"errors"
	"workspace/db"
	"workspace/models"

	"gorm.io/gorm/clause"
)

func GetParticipantes() ([]models.Participantes, error) {
	var participantes []models.Participantes

	err := db.DB.Preload(clause.Associations).Find(&participantes).Error

	return participantes, err
}

// This is not the best way to do this as it would be better to just make one query to the Formaliza table using the DNI
// Im just making it this way to practice with associations
func GetRervasOfParticipante(dni string) ([]models.Formaliza, error) {
	var reservas []models.Formaliza

	var participante models.Participantes

	res := db.DB.Where("DNI = ?", dni).Find(&participante)

	if res.RowsAffected == 0 {
		return nil, errors.New("DNI does not exist in the Data Base")
	}

	db.DB.Model(&participante).Association("Formalizaciones").Find(&reservas)

	return reservas, nil
}
