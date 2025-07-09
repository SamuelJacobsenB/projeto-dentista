package services

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/utils"
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

func (service *PatientService) Update(patient *entities.Patient, id uint) error {
	if err := service.repo.Update(patient, id); err != nil {
		return err
	}

	return nil
}

func (service *PatientService) UploadImage(profileImage *multipart.FileHeader, id uint) error {
	imageExtension := strings.TrimPrefix(filepath.Ext(profileImage.Filename), ".")

	if err := utils.ValidImageExtension(imageExtension); err != nil {
		return err
	}

	const maxSize = 3 * 1024 * 1024
	if profileImage.Size > int64(maxSize) {
		return errors.New("imagem deve ter no máximo 3MB")
	}

	if _, err := service.repo.FindByID(id); err != nil {
		return errors.New("paciente não encontrado")
	}

	if err := service.repo.UploadImage(imageExtension, id); err != nil {
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
