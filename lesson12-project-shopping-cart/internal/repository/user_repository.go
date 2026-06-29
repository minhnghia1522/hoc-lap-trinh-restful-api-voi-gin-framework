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
func (repo *userRepository) CountUsers(ctx context.Context, arg sqlc.CountUsersParams) (int64, error) {
	panic("unimplemented")
}

// CreateUser implements [IUserRepository].
func (repo *userRepository) CreateUser(ctx context.Context, userParam sqlc.CreateUserParams) (sqlc.User, error) {
	userCreated, err := repo.db.CreateUser(ctx, userParam)
	if err != nil {
		return sqlc.User{}, err
	}
	return userCreated, nil
}

// GetUser implements [IUserRepository].
func (repo *userRepository) GetUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements [IUserRepository].
func (repo *userRepository) GetUserByEmail(ctx context.Context, userEmail string) (sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserCreatedAtAsc implements [IUserRepository].
func (repo *userRepository) ListUsersUserCreatedAtAsc(ctx context.Context, arg sqlc.ListUsersUserCreatedAtAscParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserCreatedAtDesc implements [IUserRepository].
func (repo *userRepository) ListUsersUserCreatedAtDesc(ctx context.Context, arg sqlc.ListUsersUserCreatedAtDescParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserIdAsc implements [IUserRepository].
func (repo *userRepository) ListUsersUserIdAsc(ctx context.Context, arg sqlc.ListUsersUserIdAscParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// ListUsersUserIdDesc implements [IUserRepository].
func (repo *userRepository) ListUsersUserIdDesc(ctx context.Context, arg sqlc.ListUsersUserIdDescParams) ([]sqlc.User, error) {
	panic("unimplemented")
}

// RestoreUser implements [IUserRepository].
func (repo *userRepository) RestoreUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// SoftDeleteUser implements [IUserRepository].
func (repo *userRepository) SoftDeleteUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// TrashUser implements [IUserRepository].
func (repo *userRepository) TrashUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	panic("unimplemented")
}

// UpdatePassword implements [IUserRepository].
func (repo *userRepository) UpdatePassword(ctx context.Context, arg sqlc.UpdatePasswordParams) (sqlc.User, error) {
	panic("unimplemented")
}

// UpdateUser implements [IUserRepository].
func (repo *userRepository) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	panic("unimplemented")
}
