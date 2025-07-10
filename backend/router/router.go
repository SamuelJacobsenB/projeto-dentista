package router

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/bootstrap"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/db"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoter() *gin.Engine {
	router := gin.Default()

	router.Static("/uploads/patients", "./uploads/patients")

	api := router.Group("/api")
	v1 := api.Group("/v1")

	authController := bootstrap.InitAuthModule(db.DB)
	RegisterAuthRoutes(v1.Group("/auth"), authController)

	userController := bootstrap.InitUserModule(db.DB)
	userGroup := v1.Group("/users")
	userGroup.Use(middlewares.AuthMiddleware(nil))
	RegisterUserRoutes(userGroup, userController)

	patientController := bootstrap.InitPatientModule(db.DB)
	patientGroup := v1.Group("/patients")
	patientGroup.Use(middlewares.AuthMiddleware(nil))
	RegisterPatientRoutes(patientGroup, patientController)

	appointmentController := bootstrap.InitAppointmentModule(db.DB)
	appointmentGroup := v1.Group("/appointments")
	appointmentGroup.Use(middlewares.AuthMiddleware(nil))
	RegisterAppointmentRoutes(appointmentGroup, appointmentController)

	bootstrap.InitScheduler(db.DB)

	return router
}
