package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/middlewares"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/types"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(routerGroup *gin.RouterGroup, controller *controllers.AuthController) {
	routerGroup.POST("/login", controller.Login)
	routerGroup.GET("/verify/user", middlewares.AuthMiddleware(nil), controller.VerifyUser)
	routerGroup.GET("/verify/admin", middlewares.AuthMiddleware([]types.Role{types.RoleUser, types.RoleAdmin}), controller.VerifyUser)
}
