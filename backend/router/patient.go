package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPatientRoutes(routerGroup *gin.RouterGroup, controller *controllers.PatientController) {
	routerGroup.GET("/", controller.FindPagenedByName)
	routerGroup.GET("/:id", controller.FindByID)
	routerGroup.POST("/", controller.Create)
	routerGroup.PATCH("/:id", controller.Update)
	routerGroup.DELETE("/:id", controller.Delete)
}
