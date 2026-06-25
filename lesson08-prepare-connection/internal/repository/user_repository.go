package repository

import (
	"lesson08-prepare-connection/internal/models"
	"log"
)

type SQLUserRepository struct {
}

func NewSQLUserRepository() UserRepository {
	return &SQLUserRepository{}
}

// Create implements [UserRepository].
func (s *SQLUserRepository) Create(user *models.User) {
	log.Println("unimplemented")
}

// FindById implements [UserRepository].
func (s *SQLUserRepository) FindById(id int) {
	log.Println("unimplemented")
}
