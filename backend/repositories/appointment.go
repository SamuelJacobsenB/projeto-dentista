package repositories

import (
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{db}
}

func (repo *AppointmentRepository) FindAll() ([]entities.Appointment, error) {
	var appointments []entities.Appointment

	if err := repo.db.Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

func (repo *AppointmentRepository) FindOfToday() ([]entities.Appointment, error) {
	start := time.Now().Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	var appointments []entities.Appointment
	if err := repo.db.Where("start_time >= ? AND start_time < ?", start, end).Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

func (repo *AppointmentRepository) FindByID(id uint) (*entities.Appointment, error) {
	var appointment entities.Appointment

	if err := repo.db.First(&appointment, id).Error; err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (repo *AppointmentRepository) Create(appointment *entities.Appointment) error {
	if err := repo.db.Create(appointment).Error; err != nil {
		return err
	}

	return nil
}

func (repo *AppointmentRepository) Update(appointment *entities.Appointment, id uint) error {
	existing, err := repo.FindByID(id)
	if err != nil {
		return err
	}

	if err := repo.db.Model(&existing).Updates(appointment).Error; err != nil {
		return err
	}

	return nil
}

func (repo *AppointmentRepository) Delete(id uint) error {
	if err := repo.db.Delete(&entities.Appointment{}, id).Error; err != nil {
		return err
	}

	return nil
}
