package entities

import "github.com/SamuelJacobsenB/projeto-dentista/dtos/response"

type Patient struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Cpf       int    `json:"cpf" gorm:"not null;unique"`
	Phone     int    `json:"phone" gorm:"not null"`
	Address   string `json:"address" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"not null"`
	UpdatedAt string `json:"updated_at" gorm:"not null"`
}

func (patient *Patient) toResponseDTO() response.PatientDTO {
	return response.PatientDTO{
		ID:        patient.ID,
		Name:      patient.Name,
		Email:     patient.Email,
		Cpf:       patient.Cpf,
		Phone:     patient.Phone,
		Address:   patient.Address,
		CreatedAt: patient.CreatedAt,
		UpdatedAt: patient.UpdatedAt,
	}
}
