package bootstrap

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"gorm.io/gorm"
)

func InitPatientModule(db *gorm.DB) *controllers.PatientController {
	repo := repositories.NewPatientRepository(db)
	service := services.NewPatientService(repo)
	controller := controllers.NewPatientController(service)

	return controller
}
