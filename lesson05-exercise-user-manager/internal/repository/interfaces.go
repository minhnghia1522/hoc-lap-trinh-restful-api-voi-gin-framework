package repository

import "user-management-api/internal/models"

type InMemoryUserRepository interface {
	SelectByCondition()
	SelectByUUID(uuid string) (models.User, bool)
	SelectByEmail(email string) (models.User, bool)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(uuid string) error
}
