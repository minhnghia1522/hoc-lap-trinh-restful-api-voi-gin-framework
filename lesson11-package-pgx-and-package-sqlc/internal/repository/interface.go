package repository

import (
	"context"
	"lesson08-prepare-connection/internal/db/sqlc"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, userParam sqlc.CreateUserParams) (sqlc.User, error)
	FindByUUID(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
}
