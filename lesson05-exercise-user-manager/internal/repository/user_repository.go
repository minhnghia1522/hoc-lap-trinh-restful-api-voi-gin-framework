package repository

import "user-management-api/internal/models"

type inMemoryUserRepository struct {
	users []models.User
}

func NewUserRepository() InMemoryUserRepository {
	return &inMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

func (repo *inMemoryUserRepository) SelectByCondition() {

}

func (repo *inMemoryUserRepository) SelectByUUID(uuid string) {

}

func (repo *inMemoryUserRepository) CreateUser() {

}

func (repo *inMemoryUserRepository) UpdateUser() {

}

func (repo *inMemoryUserRepository) DeleteUser(uuid string) {

}
