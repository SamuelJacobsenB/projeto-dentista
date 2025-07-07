package router

import "github.com/gin-gonic/gin"

func SetupRoter() *gin.Engine {
	router := gin.Default()

	RegisterUserRoutes(router.Group("/users"))

	return router
}
