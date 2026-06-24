package service

import (
	"user-management-api/internal/models"
)

type UserService interface {
	Search()
	FindUserByUUID(uuid string) (models.User, error)
	CreateUser(userModel models.User) (models.User, error)
	UpdateUser(uuid string, userModel models.User) (models.User, error)
	DeleteUser(uuid string) error
}
