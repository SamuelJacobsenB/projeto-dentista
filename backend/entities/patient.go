package entities

import (
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/response"
)

type Patient struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Cpf       string    `json:"cpf" gorm:"not null;unique"`
	Phone     string    `json:"phone" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Unpaid    float64   `json:"unpaid" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (patient *Patient) ToResponseDTO() *response.PatientDTO {
	return &response.PatientDTO{
		ID:        patient.ID,
		Name:      patient.Name,
		Email:     patient.Email,
		Cpf:       patient.Cpf,
		Phone:     patient.Phone,
		Address:   patient.Address,
		Unpaid:    patient.Unpaid,
		CreatedAt: patient.CreatedAt,
		UpdatedAt: patient.UpdatedAt,
	}
}
