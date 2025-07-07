package repositories

import (
	"github.com/SamuelJacobsenB/projeto-dentista/backend/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) FindByID(id uint) (*entities.User, error) {
	var user entities.User

	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Create(user *entities.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Promote(id uint) error {
	if err := repo.db.Model(&entities.User{}).Where("id = ?", id).Update("roles", []string{"user, admin"}).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Delete(id uint) error {
	if err := repo.db.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
