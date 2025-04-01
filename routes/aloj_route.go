package routes

import (
	"workspace/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAlojamientoRoutes(router *gin.Engine) {

	api := router.Group("/alojamientos")
	{
		api.GET("/list", handlers.GetAlojamientos)
		api.GET("/:id", handlers.GetALojamientoById)
		api.POST("/create", handlers.CreateAlojamiento)
	}
}
