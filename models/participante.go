package models

import (
	"time"
)

type Participantes struct {
	DNI             string     `gorm:"primaryKey;column:dni;size:9"`
	Nombre          string     `gorm:"column:nombre;size:50;not null"`
	Apellido        string     `gorm:"column:apellido;size:50;not null"`
	Ciudad          *string    `gorm:"column:ciudad;size:50"`
	FechaNacimiento *time.Time `gorm:"column:fechanacimiento"`
	Telefono        *string    `gorm:"column:telefono;size:9"`

	Formalizaciones []Formaliza `gorm:"foreignKey:DNI_Cliente; references:DNI"`
}
