package controller

import (
	"net/http"

	"github.com/Parzeval8/proyecto-api-go/config"
	"github.com/Parzeval8/proyecto-api-go/models"
	"github.com/gin-gonic/gin"
)

// se usa el gin.context para capturar la peticion del cliente
func GetPruebas(c *gin.Context) {
	pruebas := []models.Prueba{}
	//se buscan los datos de la tabla
	config.DB.Find(&pruebas)
	//se responde al cliente la peticion ok y el dato tipo json
	c.IndentedJSON(http.StatusOK, &pruebas)
}
func PostPruebas(c *gin.Context) {
	var prueba models.Prueba
	//asigna el valor recibido, y se crea en la tabla
	c.BindJSON(&prueba)
	config.DB.Create(&prueba)
	//se responde al cliente la peticion ok y el dato tipo json
	c.IndentedJSON(http.StatusCreated, &prueba)
}
func DeletePruebas(c *gin.Context) {
	var prueba models.Prueba
	//se busca en la base de datos y se elimina
	config.DB.Where("id = ?", c.Param("id")).Delete(&prueba)
	//se responde al cliente la peticion ok y el dato tipo json
	c.IndentedJSON(http.StatusOK, &prueba)
}
func PutPruebas(c *gin.Context) {
	var prueba models.Prueba
	//se busca en la base de datos y se actualiza
	config.DB.Where("id = ?", c.Param("id")).First(&prueba)
	c.BindJSON(&prueba)
	config.DB.Save(&prueba)
	//se responde al cliente la peticion ok y el dato tipo json
	c.IndentedJSON(http.StatusOK, prueba)
}
