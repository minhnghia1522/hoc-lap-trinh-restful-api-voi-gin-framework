package repository

import (
	"database/sql"
	"fmt"
	"lesson08-prepare-connection/internal/models"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

// Create implements [UserRepository].
func (s *SQLUserRepository) Create(user *models.User) error {
	sqlRow := s.db.QueryRow("INSERT into users (name, email) VALUES ($1, $2) RETURNING user_id", user.Name, user.Email)
	err := sqlRow.Scan(&user.Id)
	if err != nil {
		return fmt.Errorf("failed create user: %w", err)
	}
	return nil
}

// FindById implements [UserRepository].
func (s *SQLUserRepository) FindById(id int) (models.User, error) {
	sqlRow := s.db.QueryRow("SELECT * FROM users WHERE user_id = $1", id)
	var user models.User
	err := sqlRow.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}
