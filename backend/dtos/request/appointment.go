package request

import (
	"errors"
	"strings"
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
)

type AppointmentDTO struct {
	Attendant   string    `json:"attendant"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`

	PatientID uint `json:"patient_id"`
}

func (appointmentDTO *AppointmentDTO) Validate() error {
	appointmentDTO.Attendant = strings.TrimSpace(appointmentDTO.Attendant)
	appointmentDTO.Description = strings.TrimSpace(appointmentDTO.Description)

	if appointmentDTO.Attendant == "" {
		return errors.New("atendente é obrigatório")
	}

	if len(appointmentDTO.Attendant) > 50 {
		return errors.New("nome do atendente deve ter no máximo 50 caracteres")
	}

	if appointmentDTO.StartTime.IsZero() {
		return errors.New("horário de início é obrigatório")
	}

	if appointmentDTO.EndTime.IsZero() {
		return errors.New("horário de término é obrigatório")
	}

	if appointmentDTO.StartTime.After(appointmentDTO.EndTime) {
		return errors.New("horário de início não pode ser posterior ao horário de término")
	}

	if appointmentDTO.PatientID == 0 {
		return errors.New("id do paciente é obrigatório")
	}

	return nil
}

func (appointmentDTO *AppointmentDTO) ToEntity() *entities.Appointment {
	return &entities.Appointment{
		Attendant:   appointmentDTO.Attendant,
		Description: appointmentDTO.Description,
		StartTime:   appointmentDTO.StartTime,
		EndTime:     appointmentDTO.EndTime,
		PatientID:   appointmentDTO.PatientID,
	}
}
