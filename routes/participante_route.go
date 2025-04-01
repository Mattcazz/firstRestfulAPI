package routes

import (
	"workspace/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpParticipanteRoutes(router *gin.Engine) {
	api := router.Group("/participante")
	{
		api.GET("/list", handlers.GetParticipantes)
		api.GET("/:dni/reservas", handlers.GetRervasOfParticipante)
	}

}
