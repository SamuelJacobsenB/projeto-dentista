package bootstrap

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/controllers"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"gorm.io/gorm"
)

func InitAuthModule(db *gorm.DB) *controllers.AuthController {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userService)
	return controllers.NewAuthController(authService)
}
