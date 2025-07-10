package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAppointmentRoutes(routerGroup *gin.RouterGroup, controller *controllers.AppointmentController) {
	routerGroup.GET("/", controller.FindAll)
	routerGroup.GET("/today", controller.FindOfToday)
	routerGroup.GET("/:id", controller.FindByID)
	routerGroup.POST("/", controller.Create)
	routerGroup.PATCH("/:id", controller.Update)
	routerGroup.DELETE("/:id", controller.Delete)
}
