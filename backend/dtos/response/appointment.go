package response

type AppointmentDTO struct {
	ID          uint        `json:"id"`
	Attendant   string      `json:"attendant"`
	Description string      `json:"description"`
	StartTime   string      `json:"start_time"`
	EndTime     string      `json:"end_time"`
	Reminder    bool        `json:"reminder"`
	PatientID   uint        `json:"patient_id"`
	Patient     *PatientDTO `json:"patient,omitempty"`
}
