package response

import "time"

type PatientDTO struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Cpf            string    `json:"cpf"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Unpaid         float64   `json:"unpaid"`
	ImageExtension string    `json:"image,omitempty"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Appointments []AppointmentDTO `json:"appointments,omitempty"`
}
