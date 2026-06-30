package repository

import (
	"context"
	"fmt"
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

// GetUserForUpdateNoWait implements [UserRepository].
func (repo *userRepository) GetUserForUpdateNoWait(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	user, err := repo.q.GetUserForUpdateNoWait(ctx, userUuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

// CountUsers implements [IUserRepository].
func (repo *userRepository) CountUsers(ctx context.Context, arg sqlc.CountUsersParams) (int64, error) {
	return repo.q.CountUsers(ctx, arg)
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
	user, err := repo.q.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
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
	user, err := repo.q.RestoreUser(ctx, userUuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

// SoftDeleteUser implements [IUserRepository].
func (repo *userRepository) SoftDeleteUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	user, err := repo.q.SoftDeleteUser(ctx, userUuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

// TrashUser implements [IUserRepository].
func (repo *userRepository) TrashUser(ctx context.Context, userUuid uuid.UUID) (sqlc.User, error) {
	user, err := repo.q.TrashUser(ctx, userUuid)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

// UpdatePassword implements [IUserRepository].
func (repo *userRepository) UpdatePassword(ctx context.Context, arg sqlc.UpdatePasswordParams) (sqlc.User, error) {
	user, err := repo.q.UpdatePassword(ctx, arg)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

// UpdateUser implements [IUserRepository].
func (repo *userRepository) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	user, err := repo.q.UpdateUser(ctx, arg)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
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

// GetAll implements [UserRepository].
func (repo *userRepository) GetAll(ctx context.Context, search string, orderBy string, sort string, limit int32, offset int32) ([]sqlc.User, error) {
	var (
		users []sqlc.User
		err   error
	)

	switch {
	case orderBy == "user_id" && sort == "asc":
		users, err = repo.q.ListUsersUserIdAsc(ctx, sqlc.ListUsersUserIdAscParams{
			Limit:  limit,
			Offset: offset,
			Search: search,
		})
	case orderBy == "user_id" && sort == "desc":
		users, err = repo.q.ListUsersUserIdDesc(ctx, sqlc.ListUsersUserIdDescParams{
			Limit:  limit,
			Offset: offset,
			Search: search,
		})
	case orderBy == "user_created_at" && sort == "asc":
		users, err = repo.q.ListUsersUserCreatedAtAsc(ctx, sqlc.ListUsersUserCreatedAtAscParams{
			Limit:  limit,
			Offset: offset,
			Search: search,
		})
	case orderBy == "user_created_at" && sort == "desc":
		users, err = repo.q.ListUsersUserCreatedAtDesc(ctx, sqlc.ListUsersUserCreatedAtDescParams{
			Limit:  limit,
			Offset: offset,
			Search: search,
		})
	}

	if err != nil {
		return []sqlc.User{}, err
	}

	return users, nil
}

// GetAllV2 implements [UserRepository].
func (repo *userRepository) GetAllV2(ctx context.Context, search string, orderBy string, sort string, limit int32, offset int32, deleted bool) ([]sqlc.User, error) {
	query := `SELECT *
		FROM users
		WHERE (
			$1::TEXT IS NULL
			OR $1::TEXT = ''
			OR user_email ILIKE '%' || $1 || '%'
			OR user_fullname ILIKE '%' || $1 || '%'
		)`

	if deleted {
		query += " AND user_deleted_at IS NOT NULL"
	} else {
		query += " AND user_deleted_at IS NULL"
	}

	order := "ASC"
	if sort == "desc" {
		order = "DESC"
	}

	switch orderBy {
	case "user_id", "user_created_at":
		query += fmt.Sprintf(" ORDER BY %s %s", orderBy, order)
	default:
		query += " ORDER BY user_id ASC"
	}

	query += " LIMIT $2 OFFSET $3 -- name: Get All Version 2"

	rows, err := repo.pool.Query(ctx, query, search, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []sqlc.User{}
	for rows.Next() {
		var i sqlc.User
		if err := rows.Scan(
			&i.UserID,
			&i.UserUuid,
			&i.UserEmail,
			&i.UserPassword,
			&i.UserFullname,
			&i.UserAge,
			&i.UserStatus,
			&i.UserLevel,
			&i.UserDeletedAt,
			&i.UserCreatedAt,
			&i.UserUpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
