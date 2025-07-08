package services

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
)

type PatientService struct {
	repo *repositories.PatientRepository
}

func NewPatientService(repo *repositories.PatientRepository) *PatientService {
	return &PatientService{repo}
}

func (service *PatientService) FindPagenedByName(name string, limit, offset int) ([]entities.Patient, error) {
	patients, err := service.repo.FindPagenedByName(name, limit, offset)

	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (service *PatientService) FindByID(id uint) (*entities.Patient, error) {
	patient, err := service.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (service *PatientService) Create(patient *entities.Patient) error {
	if err := service.repo.Create(patient); err != nil {
		return err
	}

	return nil
}

func (service *PatientService) Update(patient *entities.Patient) error {
	if err := service.repo.Update(patient); err != nil {
		return err
	}

	return nil
}

func (service *PatientService) Delete(id uint) error {
	if err := service.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
