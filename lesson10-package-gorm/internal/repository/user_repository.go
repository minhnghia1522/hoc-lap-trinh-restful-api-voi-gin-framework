package repository

import (
	"lesson08-prepare-connection/internal/models"

	"gorm.io/gorm"
)

type SQLUserRepository struct {
	db *gorm.DB
}

func NewSQLUserRepository(db *gorm.DB) UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

// Create implements [UserRepository].
func (s *SQLUserRepository) Create(user *models.User) error {
	if err := s.db.Create(user).Error; err != nil {

		return err
	}
	return nil
}

// FindById implements [UserRepository].
func (s *SQLUserRepository) FindById(id int) (models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil

}
