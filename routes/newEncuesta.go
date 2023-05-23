package routes

import (
	"github.com/Parzeval8/proyecto-api-go/controller"
	"github.com/gin-gonic/gin"
)

func NewEncuestaRoute(router *gin.Engine) {
	// router.GET("/encuesta", controller.GetEncuestas)
	router.POST("/newencuesta", controller.PostNewEncuesta)
	router.GET("/newencuesta/:id", controller.GetNewEncuesta)
	// router.DELETE("/:id", controller.DeletePruebas)
	// router.PUT("/:id", controller.PutPruebas)

}
