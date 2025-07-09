package entities

import (
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/response"
)

type Appointment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Attendant   string    `json:"attendant" gorm:"not null;size:50"`
	Description string    `json:"description" gorm:"type:text"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	Reminder    bool      `json:"reminder" gorm:"not null;default:false"`

	PatientID uint     `json:"patient_id"`
	Patient   *Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (appointment *Appointment) ToResponseDTO() *response.AppointmentDTO {
	return &response.AppointmentDTO{
		ID:          appointment.ID,
		Attendant:   appointment.Attendant,
		Description: appointment.Description,
		StartTime:   appointment.StartTime.Format(time.RFC3339),
		EndTime:     appointment.EndTime.Format(time.RFC3339),
		Reminder:    appointment.Reminder,
		PatientID:   appointment.PatientID,
		Patient:     appointment.Patient.ToResponseDTO(),
	}
}
