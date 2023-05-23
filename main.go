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
	router.Use(CORSMiddleware())

	routes.PruebaRoute(router)
	routes.EncuestaRoute(router)
	routes.NewEncuestaRoute(router)
	router.Run("localhost:8080")
	// router.Run("localhost:9000")
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
