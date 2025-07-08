package request

import (
	"errors"
	"strings"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/types"
)

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userDTO *UserDTO) Validate() error {
	userDTO.Name = strings.TrimSpace(userDTO.Name)
	userDTO.Email = strings.TrimSpace(userDTO.Email)
	userDTO.Password = strings.TrimSpace(userDTO.Password)

	if userDTO.Name == "" {
		return errors.New("nome é obrigatório")
	}

	if len(userDTO.Name) < 3 || len(userDTO.Name) > 50 {
		return errors.New("nome deve ter entre 3 e 50 caracteres")
	}

	if userDTO.Email == "" {
		return errors.New("email é obrigatório")
	}

	if strings.Contains(userDTO.Email, " ") || !strings.Contains(userDTO.Email, "@") {
		return errors.New("email inválido")
	}

	if userDTO.Password == "" {
		return errors.New("senha é obrigatória")
	}

	if strings.Contains(userDTO.Password, " ") {
		return errors.New("senha não pode conter espaços")
	}

	if len(userDTO.Password) < 8 || len(userDTO.Password) > 15 {
		return errors.New("senha deve ter entre 8 e 15 caracteres")
	}

	return nil
}

func (userDTO *UserDTO) ToEntity() *entities.User {
	return &entities.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
		Roles:    []types.Role{types.RoleUser},
	}
}
