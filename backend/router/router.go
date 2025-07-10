package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/bootstrap"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/db"
	"github.com/gin-gonic/gin"
)

func SetupRoter() *gin.Engine {
	router := gin.Default()

	router.Static("/uploads/patients", "./uploads/patients")

	api := router.Group("/api")
	v1 := api.Group("/v1")

	userController := bootstrap.InitUserModule(db.DB)
	RegisterUserRoutes(v1.Group("/users"), userController)

	patientController := bootstrap.InitPatientModule(db.DB)
	RegisterPatientRoutes(v1.Group("/patients"), patientController)

	appointmentController := bootstrap.InitAppointmentModule(db.DB)
	RegisterAppointmentRoutes(v1.Group("/appointments"), appointmentController)

	return router
}
