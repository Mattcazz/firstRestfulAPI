package db

import (
	"fmt"
	"workspace/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(&models.Alojamientos{}, &models.Reservas{}, &models.Participantes{}, &models.Formaliza{})

	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully")
	}

}
