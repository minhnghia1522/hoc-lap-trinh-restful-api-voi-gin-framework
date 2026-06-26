package repository

import (
	"context"
	"user-management-api/internal/db/sqlc"

	"github.com/google/uuid"
)

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{
		db: db,
	}
}

// CountUsers implements [IUserRepository].
func (u *userRepository) CountUsers(ctx context.Context, arg sqlc.CountUsersParams) (int64, error) {
	panic("unimplemented")
}

// CreateUser implements [IUserRepository].
func (u *userRepository) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	panic("unimplemented")
}

// GetUser implements [IUserRepository].
func (u *userRepository) GetUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements [IUserRepository].
func (u *userRepository) GetUserByEmail(ctx context.Context, userEmail string) (sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserCreatedAtAsc implements [IUserRepository].
func (u *userRepository) ListUsersUserCreatedAtAsc(ctx context.Context, arg sqlc.ListUsersUserCreatedAtAscParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserCreatedAtDesc implements [IUserRepository].
func (u *userRepository) ListUsersUserCreatedAtDesc(ctx context.Context, arg sqlc.ListUsersUserCreatedAtDescParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserIdAsc implements [IUserRepository].
func (u *userRepository) ListUsersUserIdAsc(ctx context.Context, arg sqlc.ListUsersUserIdAscParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserIdDesc implements [IUserRepository].
func (u *userRepository) ListUsersUserIdDesc(ctx context.Context, arg sqlc.ListUsersUserIdDescParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// RestoreUser implements [IUserRepository].
func (u *userRepository) RestoreUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// SoftDeleteUser implements [IUserRepository].
func (u *userRepository) SoftDeleteUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// TrashUser implements [IUserRepository].
func (u *userRepository) TrashUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// UpdatePassword implements [IUserRepository].
func (u *userRepository) UpdatePassword(ctx context.Context, arg sqlc.UpdatePasswordParams) (sqlc.User, error) {
	panic("unimplemented")
}

// UpdateUser implements [IUserRepository].
func (u *userRepository) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	panic("unimplemented")
}
