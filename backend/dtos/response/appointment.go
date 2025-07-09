package response

import "time"

type AppointmentDTO struct {
	ID          uint        `json:"id"`
	Attendant   string      `json:"attendant"`
	Description string      `json:"description"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
	Reminder    bool        `json:"reminder"`
	PatientID   uint        `json:"patient_id"`
	Patient     *PatientDTO `json:"patient,omitempty"`
}
