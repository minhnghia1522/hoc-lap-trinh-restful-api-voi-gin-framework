package repository

import (
	"context"
	"fmt"
	"lesson08-prepare-connection/internal/db/sqlc"

	"github.com/google/uuid"
)

type SQLUserRepository struct {
	db sqlc.Querier
}

func NewSQLUserRepository(db sqlc.Querier) UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

// Create implements [UserRepository].
func (s *SQLUserRepository) Create(ctx context.Context, userParam sqlc.CreateUserParams) (sqlc.User, error) {
	user, err := s.db.CreateUser(ctx, userParam)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

// FindById implements [UserRepository].
func (s *SQLUserRepository) FindByUUID(ctx context.Context, uuid uuid.UUID) (sqlc.User, error) {
	user, err := s.db.GetUser(ctx, uuid)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil

}
