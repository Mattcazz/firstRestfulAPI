package models

type Formaliza struct {
	IdRes       uint   `gorm:"column:idreserva"`
	DNI_Cliente string `gorm:"column:dni"`
}

func (Formaliza) TableName() string {
	return "formaliza"
}
