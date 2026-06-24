package service

import (
	"user-management-api/internal/models"
)

type UserService interface {
	Search(search string, page, limit int) []models.User
	FindUserByUUID(uuid string) (models.User, error)
	CreateUser(userModel models.User) (models.User, error)
	UpdateUser(uuid string, userModel models.User) (models.User, error)
	DeleteUser(uuid string) error
}
