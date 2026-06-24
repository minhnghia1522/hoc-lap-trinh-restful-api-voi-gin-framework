package service

import (
	"user-management-api/internal/repository"
)

type userService struct {
	repo repository.InMemoryUserRepository
}

func NewUserService(repo repository.InMemoryUserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (service *userService) Search() {
	service.repo.SelectByCondition()
}

func (service *userService) FindUserByUUID(uuid string) {
	service.repo.SelectByUUID(uuid)
}

func (service *userService) CreateUser() {
	service.repo.CreateUser()
}

func (service *userService) UpdateUser() {
	service.repo.UpdateUser()
}

func (service *userService) DeleteUser(uuid string) {
	service.repo.DeleteUser(uuid)
}
