package service

import (
	"user-management-api/internal/dto"
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"

	"github.com/google/uuid"
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

func (service *userService) FindUserByUUID(uuid string) (models.User, error) {
	model, exist := service.repo.SelectByUUID(uuid)
	if !exist {
		return models.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
	}
	return model, nil
}

func (service *userService) CreateUser(userRequest dto.CreateUserInput) (models.User, error) {
	userModel := userRequest.MapCreateInputToModel()

	if _, exist := service.repo.SelectByEmail(userModel.Email); exist {
		return models.User{}, utils.NewError("email already exist", utils.ErrCodeConflict)
	}

	uuid := uuid.New().String()
	userModel.UUID = uuid
	service.repo.CreateUser(userModel)
	return userModel, nil
}

func (service *userService) UpdateUser(userRequest dto.UpdateUserInput) {
	userModel := userRequest.MapUpdateInputToModel()
	service.repo.UpdateUser(userModel)
}

func (service *userService) DeleteUser(uuid string) {
	service.repo.DeleteUser(uuid)
}
