package main

import (
	"github.com/Parzeval8/proyecto-api-go/config"
	"github.com/Parzeval8/proyecto-api-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect()
	// se crea router para poder inicializar las rutas
	router := gin.New()

	routes.PruebaRoute(router)
	routes.EncuestaRoute(router)
	router.Run("localhost:8080")
}
