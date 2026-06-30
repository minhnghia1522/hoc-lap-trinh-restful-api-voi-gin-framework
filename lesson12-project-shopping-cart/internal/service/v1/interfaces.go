package v1service

import (
	"time"
	"user-management-api/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Search(search string, page, limit int) []sqlc.User
	FindUserByUUID(ctx *gin.Context, uuid string) (sqlc.User, error)
	CreateUser(ctx *gin.Context, userParam sqlc.CreateUserParams) (sqlc.User, error)
	UpdateUser(ctx *gin.Context, uuid string, updatedAt time.Time, params sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(uuid string) error
}
