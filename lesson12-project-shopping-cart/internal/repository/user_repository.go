package repository

import (
	"context"
	"user-management-api/internal/db/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepository{
		pool: pool,
		q:    sqlc.New(pool),
	}
}

// GetUserForUpdate implements [UserRepository].
func (repo *userRepository) GetUserForUpdate(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	user, err := repo.q.GetUserForUpdate(ctx, userUuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

// CountUsers implements [IUserRepository].
func (repo *userRepository) CountUsers(ctx context.Context, arg sqlc.CountUsersParams) (int64, error) {
	panic("unimplemented")
}

// CreateUser implements [IUserRepository].
func (repo *userRepository) CreateUser(ctx context.Context, userParam sqlc.CreateUserParams) (sqlc.User, error) {
	userCreated, err := repo.q.CreateUser(ctx, userParam)
	if err != nil {
		return sqlc.User{}, err
	}
	return userCreated, nil
}

// GetUser implements [IUserRepository].
func (repo *userRepository) GetUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	user, err := repo.q.GetUser(ctx, userUuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
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

func (repo *userRepository) ExecTx(
	ctx context.Context,
	fn func(*sqlc.Queries) error,
) error {
	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	q := repo.q.WithTx(tx)

	if err := fn(q); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
