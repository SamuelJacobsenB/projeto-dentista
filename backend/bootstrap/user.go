package bootstrap

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"gorm.io/gorm"
)

func InitUserModule(db *gorm.DB) *controllers.UserController {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	return controller
}
