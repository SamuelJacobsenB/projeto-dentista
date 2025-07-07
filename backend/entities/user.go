package entities

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/types"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/response"
)

type User struct {
	ID       uint         `json:"id" gorm:"primaryKey"`
	Name     string       `json:"name" gorm:"not null;size:50"`
	Email    string       `json:"email" gorm:"unique"`
	Password string       `json:"-" gorm:"not null;size:15"`
	Roles    []types.Role `json:"roles" gorm:"many2many:user_roles;"`
}

func (user *User) toResponseDTO() response.UserDTO {
	return response.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Roles: user.Roles,
	}
}
