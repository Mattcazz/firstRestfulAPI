package handlers

import (
	"net/http"
	"strconv"
	"workspace/models"
	"workspace/services"

	"github.com/gin-gonic/gin"
)

func GetAlojamientos(c *gin.Context) {

	alojamientos, err := services.GetAlojamientos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch alojamientos"})
		return
	}

	c.IndentedJSON(http.StatusOK, alojamientos)
}

func GetALojamientoById(c *gin.Context) {
	id_int, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id number not valid"})
		return
	}

	alojamiento, err := services.GetALojamientoById(uint(id_int))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get alojamiento"})
		return
	}

	c.IndentedJSON(http.StatusOK, alojamiento)
}

func CreateAlojamiento(c *gin.Context) {

	var alojamiento models.Alojamientos

	if err := c.BindJSON(&alojamiento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateAlojamiento(&alojamiento)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, alojamiento)

}
