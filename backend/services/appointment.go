package services

import (
	"fmt"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/utils"
)

type AppointmentService struct {
	repository *repositories.AppointmentRepository
}

func NewAppointmentService(repository *repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{repository}
}

func (service *AppointmentService) FindAll() ([]entities.Appointment, error) {
	appointments, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (service *AppointmentService) FindOfToday() ([]entities.Appointment, error) {
	appointments, err := service.repository.FindOfToday()
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (service *AppointmentService) FindByID(id uint) (*entities.Appointment, error) {
	appointment, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (service *AppointmentService) Create(appointment *entities.Appointment) error {
	if err := service.repository.Create(appointment); err != nil {
		return err
	}

	return nil
}

func (service *AppointmentService) Update(appointment *entities.Appointment, id uint) error {
	if err := service.repository.Update(appointment, id); err != nil {
		return err
	}

	return nil
}

func (service *AppointmentService) SendReminderEmail() error {
	appointments, err := service.FindOfToday()
	if err != nil {
		return err
	}

	for _, appointment := range appointments {
		if !appointment.Reminder {
			if err := utils.SendEmail(appointment.Patient.Email, "Lembrete de consulta", utils.GenerateBodyText(appointment.StartTime, appointment.EndTime)); err != nil {
				fmt.Printf("houve um erro ao enviar um email para %s", appointment.Patient.Email)
				continue
			} else {
				if err := service.repository.UpdateReminder(appointment.ID); err != nil {
					fmt.Printf("houve um erro ao autalizar o campo de lembrete no banco de dados para a consulto de id: %d", appointment.ID)
				}
			}
		}
	}

	return nil
}

func (service *AppointmentService) DeleteExpired() error {
	if err := service.repository.DeleteExpired(); err != nil {
		return err
	}

	return nil
}

func (service *AppointmentService) Delete(id uint) error {
	if err := service.repository.Delete(id); err != nil {
		return err
	}

	return nil
}
