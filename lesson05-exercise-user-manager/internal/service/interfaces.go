package service

import (
	"user-management-api/internal/dto"
	"user-management-api/internal/models"
)

type UserService interface {
	Search()
	FindUserByUUID(uuid string) (models.User, error)
	CreateUser(userRequest dto.CreateUserInput) (models.User, error)
	UpdateUser(userRequest dto.UpdateUserInput)
	DeleteUser(uuid string)
}
