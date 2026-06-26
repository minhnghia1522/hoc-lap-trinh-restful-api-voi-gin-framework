package v1service

import (
	"user-management-api/internal/db/sqlc"
	"user-management-api/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// CreateUser implements [UserService].
func (u *userService) CreateUser(userModel sqlc.User) (sqlc.User, error) {
	panic("unimplemented")
}

// DeleteUser implements [UserService].
func (u *userService) DeleteUser(uuid string) error {
	panic("unimplemented")
}

// FindUserByUUID implements [UserService].
func (u *userService) FindUserByUUID(uuid string) (sqlc.User, error) {
	panic("unimplemented")
}

// Search implements [UserService].
func (u *userService) Search(search string, page int, limit int) []sqlc.User {
	panic("unimplemented")
}

// UpdateUser implements [UserService].
func (u *userService) UpdateUser(uuid string, userModel sqlc.User) (sqlc.User, error) {
	panic("unimplemented")
}
