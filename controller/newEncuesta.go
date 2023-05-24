package controller

import (
	"net/http"

	"github.com/Parzeval8/proyecto-api-go/config"
	"github.com/Parzeval8/proyecto-api-go/datamodel"
	"github.com/Parzeval8/proyecto-api-go/models"
	"github.com/gin-gonic/gin"
)

func GetNewEncuesta(c *gin.Context) {
	encuesta := models.NewEncuesta{}
	if err := config.DB.Where("ID = ?", c.Param("id")).First(&encuesta).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}
	c.IndentedJSON(http.StatusOK, &encuesta)
}
func PostNewEncuesta(c *gin.Context) {
	var encuesta models.NewEncuesta
	//asigna el valor recibido, y se crea en la tabla
	c.BindJSON(&encuesta)
	datamodel.Knn(&encuesta)
	if err := config.DB.Create(&encuesta).Error; err != nil {
		// Error al crear la encuesta en la base de datos
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la encuesta"})
		return
	}
	//se responde al cliente la peticion ok y el dato tipo json
	response := gin.H{
		"Position":   []interface{}{&encuesta.P3Edad, &encuesta.P39HabiPrac, &encuesta.P1Carrera},
		"Color":      "red",
		"P33GraMoti": &encuesta.P33GraMoti,
	}

	c.IndentedJSON(http.StatusCreated, &response)
}
