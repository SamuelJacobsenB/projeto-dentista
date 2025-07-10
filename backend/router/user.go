package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/middlewares"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/types"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(routerGroup *gin.RouterGroup, controller *controllers.UserController) {
	routerGroup.GET("/:id", controller.FindByID)
	routerGroup.POST("/", middlewares.AuthMiddleware([]types.Role{types.RoleUser, types.RoleAdmin}), controller.Create)
	routerGroup.PATCH("/:id", middlewares.AuthMiddleware([]types.Role{types.RoleUser, types.RoleAdmin}), controller.Promote)
	routerGroup.DELETE("/:id", middlewares.AuthMiddleware([]types.Role{types.RoleUser, types.RoleAdmin}), controller.Delete)
}
