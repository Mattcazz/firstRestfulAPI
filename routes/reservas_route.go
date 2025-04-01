package routes

import (
	"workspace/handlers"

	"github.com/gin-gonic/gin"
)

func SetupReservasRoutes(router *gin.Engine) {
	api := router.Group("/reservas")
	{
		api.GET("/list", handlers.GetReservas)
	}
}
