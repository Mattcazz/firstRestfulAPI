package models

import (
	"time"
)

type Reservas struct {
	IdReserva    uint      `gorm:"primaryKey;column:idreserva"`
	IdAloj       uint      `gorm:"not null;column:idalojamientos"` // Foreign Key
	FechaEntrada time.Time `gorm:"not null;column:fechaentrada"`
	FechaSalida  time.Time `gorm:"not null;column:fechasalida"`
	Precio       int       `gorm:"not null;column:precio"`

	// Relationship
	Formalizacion []Formaliza `gorm:"foreignKey:IdRes;references:IdReserva"`
}

