package repository

import "user-management-api/internal/models"

type InMemoryUserRepository interface {
	SelectByCondition()
	SelectByUUID(uuid string) (models.User, bool)
	SelectByEmail(email string) (models.User, bool)
	Insert(user models.User) error
	Update(user models.User) error
	Delete(uuid string) error
}
