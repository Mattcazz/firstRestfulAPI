package services

import (
	"errors"
	"workspace/models"
	"workspace/repositories"
)

func GetParticipantes() ([]models.Participantes, error) {
	return repositories.GetParticipantes()
}

func GetRervasOfParticipante(dni string) ([]models.Formaliza, error) {

	if dni == "" {
		return nil, errors.New("the dni param is empty")
	}

	return repositories.GetRervasOfParticipante(dni)
}
