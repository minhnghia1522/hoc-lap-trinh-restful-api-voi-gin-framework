package repository

import (
	"lesson08-prepare-connection/internal/models"
	"log"

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
func (s *SQLUserRepository) Create(user *models.User) {
	log.Println("unimplemented")
}

// FindById implements [UserRepository].
func (s *SQLUserRepository) FindById(id int) {
	log.Println("unimplemented")
}
