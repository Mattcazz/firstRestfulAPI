package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	SetupAlojamientoRoutes(router)
	SetupReservasRoutes(router)
	SetUpParticipanteRoutes(router)
}
