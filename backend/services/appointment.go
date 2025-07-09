package services

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
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

func (service *AppointmentService) Delete(id uint) error {
	if err := service.repository.Delete(id); err != nil {
		return err
	}

	return nil
}
