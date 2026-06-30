package v1service

import (
	"database/sql"
	"errors"
	"time"
	"user-management-api/internal/db/sqlc"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// CreateUser implements [UserService].
func (us *userService) CreateUser(ctx *gin.Context, userParam sqlc.CreateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()
	userParam.UserEmail = utils.NormalizeString(userParam.UserEmail)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userParam.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, utils.WrapError(err, "Failed to hash password", utils.ErrCodeInternal)
	}
	userParam.UserPassword = string(hashedPassword)
	userCreated, err := us.repo.CreateUser(context, userParam)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return sqlc.User{}, utils.NewError("email already existed", utils.ErrCodeConflict)
		}
		return sqlc.User{}, utils.WrapError(err, "Failed to insert user", utils.ErrCodeInternal)
	}
	return userCreated, nil
}

// DeleteUser implements [UserService].
func (us *userService) DeleteUser(ctx *gin.Context, uuid uuid.UUID) error {
	context := ctx.Request.Context()
	_, err := us.repo.TrashUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("User not found", utils.ErrCodeNotFound)
		}

		return utils.WrapError(err, "failed to restore user", utils.ErrCodeInternal)
	}
	return nil
}

// FindUserByUUID implements [UserService].
func (us *userService) FindUserByUUID(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()
	user, err := us.repo.GetUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found!", utils.ErrCodeNotFound)
		}
		return sqlc.User{}, utils.WrapError(err, "failed to get an user", utils.ErrCodeInternal)
	}
	return user, nil
}

// Search implements [UserService].
func (us *userService) Search(search string, page int, limit int) []sqlc.User {
	panic("unimplemented")
}

// UpdateUser implements [UserService].
func (us *userService) UpdateUser(ctx *gin.Context, uuid uuid.UUID, updatedAt time.Time, params sqlc.UpdateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()
	var updatedUser sqlc.User

	err := us.repo.ExecTx(context, func(q *sqlc.Queries) error {
		var pgErr *pgconn.PgError
		user, err := q.GetUserForUpdateNoWait(context, uuid)
		if err != nil {
			if errors.As(err, &pgErr) && pgErr.Code == "55P03" {
				return utils.NewError("Data not available", utils.ErrCodeConflict)
			}
			return err
		}

		if !user.UserUpdatedAt.Equal(updatedAt) {
			return utils.NewError("User has updated before", utils.ErrCodeConflict)
		}
		time.Sleep(5 * time.Second)
		params.UserUuid = uuid
		updatedUser, err = q.UpdateUser(ctx, params)
		return err
	})

	if err != nil {
		return sqlc.User{}, err
	}

	return updatedUser, nil

}

// RestoreUser implements [UserService].
func (us *userService) RestoreUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()
	user, err := us.repo.RestoreUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
		}

		return sqlc.User{}, utils.WrapError(err, "failed to restore user", utils.ErrCodeInternal)
	}
	return user, nil
}

// SoftDeleteUser implements [UserService].
func (us *userService) SoftDeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()
	user, err := us.repo.SoftDeleteUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
		}

		return sqlc.User{}, utils.WrapError(err, "failed to restore user", utils.ErrCodeInternal)
	}
	return user, nil
}

func (us *userService) GetAllUsers(ctx *gin.Context, search, orderBy, sort string, page, limit int32, deleted bool) ([]sqlc.User, int32, error) {
	panic("unimplemented")
}
