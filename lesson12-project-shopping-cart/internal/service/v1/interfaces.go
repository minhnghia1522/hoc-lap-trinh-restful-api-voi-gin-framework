package v1service

import (
	"time"
	"user-management-api/internal/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService interface {
	Search(search string, page, limit int) []sqlc.User
	FindUserByUUID(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error)
	CreateUser(ctx *gin.Context, userParam sqlc.CreateUserParams) (sqlc.User, error)
	UpdateUser(ctx *gin.Context, uuid uuid.UUID, updatedAt time.Time, params sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx *gin.Context, uuid uuid.UUID) error
	SoftDeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error)
	RestoreUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error)
}
