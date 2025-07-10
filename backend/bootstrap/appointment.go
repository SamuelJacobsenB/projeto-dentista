package bootstrap

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"gorm.io/gorm"
)

func InitAppointmentModule(db *gorm.DB) *controllers.AppointmentController {
	repo := repositories.NewAppointmentRepository(db)
	service := services.NewAppointmentService(repo)
	return controllers.NewAppointmentController(service)
}
