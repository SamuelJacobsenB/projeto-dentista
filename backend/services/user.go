package services

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/utils"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (service *UserService) FindByID(id uint) (*entities.User, error) {
	user, err := service.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) FindByEmail(email string) (*entities.User, error) {
	user, err := service.repo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) Create(user *entities.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := service.repo.Create(user); err != nil {
		return err
	}

	return nil
}

func (service *UserService) Promote(id uint) error {
	if err := service.repo.Promote(id); err != nil {
		return err
	}

	return nil
}

func (service *UserService) Delete(id uint) error {
	if err := service.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
