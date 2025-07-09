package request

import (
	"errors"
	"strings"
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
)

type PatientDTO struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Cpf         string    `json:"cpf"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func (patientDTO *PatientDTO) Validate() error {
	patientDTO.Name = strings.TrimSpace(patientDTO.Name)
	patientDTO.Email = strings.TrimSpace(patientDTO.Email)
	patientDTO.Address = strings.TrimSpace(patientDTO.Address)

	if patientDTO.Name == "" {
		return errors.New("nome é obrigatório")
	}

	if len(patientDTO.Name) < 3 || len(patientDTO.Name) > 100 {
		return errors.New("nome deve ter entre 3 e 100 caracteres")
	}

	if patientDTO.Email == "" {
		return errors.New("email é obrigatório")
	}

	if strings.Contains(patientDTO.Email, " ") || !strings.Contains(patientDTO.Email, "@") {
		return errors.New("email inválido")
	}

	if len(patientDTO.Cpf) != 11 {
		return errors.New("cpf inválido, deve ter 11 dígitos")
	}

	if len(patientDTO.Phone) != 10 {
		return errors.New("telefone inválido, deve ter 10 dígitos")
	}

	if patientDTO.Address == "" {
		return errors.New("endereço é obrigatório")
	}

	if len(patientDTO.Address) > 200 {
		return errors.New("endereço deve ter no máximo 200 caracteres")
	}

	if patientDTO.DateOfBirth.IsZero() {
		return errors.New("data de nascimento é obrigatória")
	}

	return nil
}

func (patientDTO *PatientDTO) ValidateUpdateDTO() error {
	patientDTO.Name = strings.TrimSpace(patientDTO.Name)
	patientDTO.Email = strings.TrimSpace(patientDTO.Email)
	patientDTO.Address = strings.TrimSpace(patientDTO.Address)

	if patientDTO.Name != "" {
		if len(patientDTO.Name) < 3 || len(patientDTO.Name) > 100 {
			return errors.New("nome deve ter entre 3 e 100 caracteres")
		}
	}

	if patientDTO.Email != "" && (strings.Contains(patientDTO.Email, " ") || !strings.Contains(patientDTO.Email, "@")) {
		return errors.New("email inválido")
	}

	if patientDTO.Cpf != "" && len(patientDTO.Cpf) != 11 {
		return errors.New("cpf inválido, deve ter 11 dígitos")
	}

	if patientDTO.Phone != "" && len(patientDTO.Phone) != 10 {
		return errors.New("telefone inválido, deve ter 10 dígitos")
	}

	if patientDTO.Address != "" && len(patientDTO.Address) > 200 {
		return errors.New("endereço deve ter no máximo 200 caracteres")
	}

	return nil
}

func (patientDTO *PatientDTO) ToEntity() *entities.Patient {
	return &entities.Patient{
		Name:        patientDTO.Name,
		Email:       patientDTO.Email,
		Cpf:         patientDTO.Cpf,
		Phone:       patientDTO.Phone,
		Address:     patientDTO.Address,
		DateOfBirth: patientDTO.DateOfBirth,
	}
}
