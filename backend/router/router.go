package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/bootstrap"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/db"
	"github.com/gin-gonic/gin"
)

func SetupRoter() *gin.Engine {
	router := gin.Default()

	userController := bootstrap.InitUserModule(db.DB)
	RegisterUserRoutes(router.Group("/users"), userController)

	patientController := bootstrap.InitPatientModule(db.DB)
	RegisterPatientRoutes(router.Group("/patients"), patientController)

	return router
}
