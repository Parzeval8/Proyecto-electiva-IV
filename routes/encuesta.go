package routes

import (
	"github.com/Parzeval8/proyecto-api-go/controller"
	"github.com/gin-gonic/gin"
)

func EncuestaRoute(router *gin.Engine) {
	router.GET("/encuestas", controller.GetEncuestas)
	router.POST("/encuestas", controller.PostEncuestas)
	router.GET("/encuesta/:id", controller.GetEncuesta)
	// router.DELETE("/:id", controller.DeletePruebas)
	// router.PUT("/:id", controller.PutPruebas)

}
