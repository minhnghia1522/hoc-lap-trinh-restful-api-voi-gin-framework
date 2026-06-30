package repository

import (
	"context"
	"user-management-api/internal/db/sqlc"

	"github.com/google/uuid"
)

type UserRepository interface {
	CountUsers(ctx context.Context, arg sqlc.CountUsersParams) (int64, error)
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error)
	GetUserForUpdate(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (sqlc.User, error)
	ListUsersUserCreatedAtAsc(ctx context.Context, arg sqlc.ListUsersUserCreatedAtAscParams) ([]sqlc.User, error)
	ListUsersUserCreatedAtDesc(ctx context.Context, arg sqlc.ListUsersUserCreatedAtDescParams) ([]sqlc.User, error)
	ListUsersUserIdAsc(ctx context.Context, arg sqlc.ListUsersUserIdAscParams) ([]sqlc.User, error)
	ListUsersUserIdDesc(ctx context.Context, arg sqlc.ListUsersUserIdDescParams) ([]sqlc.User, error)
	RestoreUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error)
	SoftDeleteUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error)
	TrashUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error)
	UpdatePassword(ctx context.Context, arg sqlc.UpdatePasswordParams) (sqlc.User, error)
	UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
	ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error
}
