package repository

import (
	"fmt"
	"slices"
	"strings"
	"user-management-api/internal/models"
	"user-management-api/internal/utils"
)

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

func (repo *inMemoryUserRepository) SelectByUUID(uuid string) (models.User, bool) {
	for _, user := range repo.users {
		if user.UUID == uuid {
			return user, true
		}
	}
	return models.User{}, false
}

func (repo *inMemoryUserRepository) SelectByEmail(email string) (models.User, bool) {
	for _, user := range repo.users {
		if match := strings.EqualFold(utils.NormalizeString(user.Email), utils.NormalizeString(email)); match {
			return user, true
		}
	}
	return models.User{}, false
}

func (repo *inMemoryUserRepository) Insert(user models.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *inMemoryUserRepository) Update(user models.User) error {
	for index, model := range repo.users {
		if model.UUID == user.UUID {
			repo.users[index] = user
			return nil
		}
	}
	return fmt.Errorf("User not found")
}

func (repo *inMemoryUserRepository) Delete(uuid string) error {
	for i, u := range repo.users {
		if u.UUID == uuid {
			repo.users = slices.Delete(repo.users, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("user not found")

}
