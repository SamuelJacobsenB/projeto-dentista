package entities

import (
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/response"
)

type Patient struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null"`
	Email          string    `json:"email" gorm:"not null"`
	Cpf            string    `json:"cpf" gorm:"not null;unique"`
	Phone          string    `json:"phone" gorm:"not null"`
	Address        string    `json:"address" gorm:"not null"`
	Unpaid         float64   `json:"unpaid" gorm:"not null;default:0"`
	ImageExtension string    `json:"image_extension,omitempty" gorm:"type:varchar(10)"`
	DateOfBirth    time.Time `json:"date_of_birth" gorm:"type:date"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null"`

	Appointments []Appointment `json:"appointments,omitempty" gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (patient *Patient) ToResponseDTO() *response.PatientDTO {
	appointments := make([]response.AppointmentDTO, len(patient.Appointments))
	for i, appointment := range patient.Appointments {
		appointments[i] = *appointment.ToResponseDTO()
	}

	return &response.PatientDTO{
		ID:             patient.ID,
		Name:           patient.Name,
		Email:          patient.Email,
		Cpf:            patient.Cpf,
		Phone:          patient.Phone,
		Address:        patient.Address,
		Unpaid:         patient.Unpaid,
		ImageExtension: patient.ImageExtension,
		DateOfBirth:    patient.DateOfBirth,
		CreatedAt:      patient.CreatedAt,
		UpdatedAt:      patient.UpdatedAt,

		Appointments: appointments,
	}
}
