package routes

import (
	"github.com/Parzeval8/proyecto-api-go/controller"
	"github.com/gin-gonic/gin"
)

func PruebaRoute(router *gin.Engine) {
	router.GET("/", controller.GetPruebas)
	router.POST("/", controller.PostPruebas)
	router.DELETE("/:id", controller.DeletePruebas)
	router.PUT("/:id", controller.PutPruebas)

}
