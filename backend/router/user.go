package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(routerGroup *gin.RouterGroup, controller *controllers.UserController) {
	routerGroup.GET("/")
}
