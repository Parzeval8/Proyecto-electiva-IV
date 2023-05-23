package controller

import (
	"net/http"

	"github.com/Parzeval8/proyecto-api-go/config"
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
	config.DB.Create(&encuesta)
	//se responde al cliente la peticion ok y el dato tipo json
	c.IndentedJSON(http.StatusCreated, &encuesta)
}