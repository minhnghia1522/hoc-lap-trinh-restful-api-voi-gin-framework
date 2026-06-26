package v1service

import (
	"user-management-api/internal/db/sqlc"
)

type UserService interface {
	Search(search string, page, limit int) []sqlc.User
	FindUserByUUID(uuid string) (sqlc.User, error)
	CreateUser(userModel sqlc.User) (sqlc.User, error)
	UpdateUser(uuid string, userModel sqlc.User) (sqlc.User, error)
	DeleteUser(uuid string) error
}
