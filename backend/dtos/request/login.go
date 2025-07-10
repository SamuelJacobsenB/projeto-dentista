package request

import (
	"errors"
	"strings"
)

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (loginDTO *LoginDTO) Validate() error {
	loginDTO.Email = strings.TrimSpace(loginDTO.Email)
	loginDTO.Password = strings.TrimSpace(loginDTO.Password)

	if loginDTO.Email == "" {
		return errors.New("email é obrigatório")
	}

	if strings.Contains(loginDTO.Email, " ") || !strings.Contains(loginDTO.Email, "@") {
		return errors.New("email inválido")
	}

	if loginDTO.Password == "" {
		return errors.New("senha é obrigatória")
	}

	if strings.Contains(loginDTO.Password, " ") {
		return errors.New("senha não pode conter espaços")
	}

	if len(loginDTO.Password) < 8 || len(loginDTO.Password) > 15 {
		return errors.New("senha deve ter entre 8 e 15 caracteres")
	}

	return nil
}
