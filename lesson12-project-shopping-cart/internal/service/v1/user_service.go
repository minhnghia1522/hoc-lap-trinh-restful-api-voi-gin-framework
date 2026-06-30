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
func (us *userService) DeleteUser(uuid string) error {
	panic("unimplemented")
}

// FindUserByUUID implements [UserService].
func (us *userService) FindUserByUUID(ctx *gin.Context, uuidPram string) (sqlc.User, error) {
	context := ctx.Request.Context()
	uuidParsed := uuid.MustParse(uuidPram)
	user, err := us.repo.GetUser(context, uuidParsed)
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
func (us *userService) UpdateUser(ctx *gin.Context, uuidParam string, updatedAt time.Time, params sqlc.UpdateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()
	var updatedUser sqlc.User
	userUuid := uuid.MustParse(uuidParam)
	err := us.repo.ExecTx(context, func(q *sqlc.Queries) error {
		var pgErr *pgconn.PgError
		user, err := q.GetUserForUpdateNoWait(context, userUuid)
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
		params.UserUuid = userUuid
		updatedUser, err = q.UpdateUser(ctx, params)
		return err
	})

	if err != nil {
		return sqlc.User{}, err
	}

	return updatedUser, nil

}
