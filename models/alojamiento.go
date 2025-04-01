package models

type Alojamientos struct {
	IdAlojamiento uint   `gorm:"primaryKey;column:idalojamiento" json:"id_alojamiento"`
	MaxPersonas   uint   `gorm:"column:maxpersonas" json:"max_personas"`
	Propietario   string `gorm:"column:propietario;not null" json:"propietario"`
	Ciudad        string `gorm:"column:ciudad;not null" json:"ciudad"`

	// Relationship
	Reservas []Reservas `gorm:"foreignKey:IdAloj;references:IdAlojamiento"`
}
