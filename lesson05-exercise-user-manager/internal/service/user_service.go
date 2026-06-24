package service

import (
	"strings"
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (service *userService) CreateUser(userModel models.User) (models.User, error) {
	if _, exist := service.repo.SelectByEmail(userModel.Email); exist {
		return models.User{}, utils.NewError("email already exist", utils.ErrCodeConflict)
	}

	uuid := uuid.New().String()
	userModel.UUID = uuid

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, utils.WrapError(err, "failed to hash password", utils.ErrCodeInternal)
	}
	userModel.Password = string(hashedPassword)

	if err := service.repo.Insert(userModel); err != nil {
		return models.User{}, utils.WrapError(err, "failed to create user", utils.ErrCodeInternal)
	}
	return userModel, nil
}

func (service *userService) UpdateUser(uuid string, userModel models.User) (models.User, error) {
	currentUser, exist := service.repo.SelectByUUID(uuid)
	if !exist {
		return models.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
	}
	if match := strings.EqualFold(utils.NormalizeString(currentUser.Email), utils.NormalizeString(userModel.Email)); !match {
		if _, exist = service.repo.SelectByEmail(userModel.Email); exist {
			return models.User{}, utils.NewError("email already exist", utils.ErrCodeConflict)
		}
	}
	currentUser.Name = userModel.Name
	currentUser.Email = userModel.Email
	currentUser.Age = userModel.Age
	currentUser.Status = userModel.Status
	currentUser.Level = userModel.Level

	if userModel.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, utils.WrapError(err, "faild to hash password", utils.ErrCodeInternal)
		}

		currentUser.Password = string(hashedPassword)
	}

	if error := service.repo.Update(currentUser); error != nil {
		return models.User{}, utils.WrapError(error, "failed to update user", utils.ErrCodeInternal)
	}
	return currentUser, nil
}

func (service *userService) DeleteUser(uuid string) error {
	_, exist := service.repo.SelectByUUID(uuid)
	if !exist {
		return utils.NewError("User not found", utils.ErrCodeNotFound)
	}

	if err := service.repo.Delete(uuid); err != nil {
		return utils.WrapError(err, "failed to delete user", utils.ErrCodeInternal)
	}
	return nil
}
