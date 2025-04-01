package handlers

import (
	"net/http"
	"workspace/services"

	"github.com/gin-gonic/gin"
)

func GetReservas(c *gin.Context) {
	reservas, err := services.GetReservas()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.IndentedJSON(http.StatusOK, reservas)
}
