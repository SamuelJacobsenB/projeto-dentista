package router

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/")
}
