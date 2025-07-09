package repositories

import (
	"errors"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db}
}

func (repo *PatientRepository) FindPagenedByName(name string, limit, offset int) ([]entities.Patient, error) {
	var patients []entities.Patient

	if err := repo.db.Where("name LIKE ?", "%"+name+"%").Limit(limit).Offset(offset).Find(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}

func (repo *PatientRepository) FindByID(id uint) (*entities.Patient, error) {
	var patient entities.Patient

	if err := repo.db.First(&patient, id).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (repo *PatientRepository) Create(patient *entities.Patient) error {
	if err := repo.db.Create(patient).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PatientRepository) Update(patient *entities.Patient, id uint) error {
	existing, err := repo.FindByID(id)
	if err != nil {
		return errors.New("paciente não encontrado")
	}

	if err := repo.db.Model(&existing).Updates(patient).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PatientRepository) UploadImage(imageExtension string, id uint) error {
	var patient entities.Patient

	if err := repo.db.First(&patient, id).Error; err != nil {
		return err
	}

	patient.ImageExtension = imageExtension

	if err := repo.db.Save(&patient).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PatientRepository) Delete(id uint) error {
	if err := repo.db.Delete(&entities.Patient{}, id).Error; err != nil {
		return err
	}

	return nil
}
