package v1service

import (
	"user-management-api/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Search(search string, page, limit int) []sqlc.User
	FindUserByUUID(uuid string) (sqlc.User, error)
	CreateUser(ctx *gin.Context, userParam sqlc.CreateUserParams) (sqlc.User, error)
	UpdateUser(uuid string, userModel sqlc.User) (sqlc.User, error)
	DeleteUser(uuid string) error
}
