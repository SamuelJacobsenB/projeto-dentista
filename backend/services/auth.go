package services

import (
	"errors"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/request"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/utils"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{userService}
}

func (service *AuthService) Login(loginDTO request.LoginDTO) (string, error) {
	user, err := service.userService.FindByEmail(loginDTO.Email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(loginDTO.Password, user.Password) {
		return "", errors.New("email ou senha incorretos, tente novamente")
	}

	token, err := utils.GenerateJwt(user.ID, user.Roles)
	if err != nil {
		return "", err
	}

	return token, nil
}
