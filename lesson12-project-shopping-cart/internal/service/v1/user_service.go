package v1service

import (
	"errors"
	"user-management-api/internal/db/sqlc"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"

	"github.com/gin-gonic/gin"
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
func (us *userService) FindUserByUUID(uuid string) (sqlc.User, error) {
	panic("unimplemented")
}

// Search implements [UserService].
func (us *userService) Search(search string, page int, limit int) []sqlc.User {
	panic("unimplemented")
}

// UpdateUser implements [UserService].
func (us *userService) UpdateUser(uuid string, userModel sqlc.User) (sqlc.User, error) {
	panic("unimplemented")
}
