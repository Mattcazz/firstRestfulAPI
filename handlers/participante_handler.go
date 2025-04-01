package handlers

import (
	"net/http"
	"workspace/services"

	"github.com/gin-gonic/gin"
)

func GetParticipantes(c *gin.Context) {
	participantes, err := services.GetParticipantes()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, participantes)
}

func GetRervasOfParticipante(c *gin.Context) {

	reservas, err := services.GetRervasOfParticipante(c.Param("dni"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, reservas)
}
