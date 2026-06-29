package v1dto

import (
	"time"
	"user-management-api/internal/db/sqlc"
	"user-management-api/internal/utils"
)

type UserDTO struct {
	UUID      string `json:"uuid"`
	Name      string `json:"full_name"`
	Email     string `json:"email_address"`
	Age       *int   `json:"age"`
	Status    string `json:"status"`
	Level     string `json:"level"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"omitempty,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

func (input *CreateUserInput) MapCreateInputToModel() sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		UserEmail:    input.Email,
		UserPassword: input.Password,
		UserFullname: input.Email,
		UserAge:      utils.ConvertToInt32Pointer(input.Age),
		UserStatus:   int32(input.Status),
		UserLevel:    int32(input.Level),
	}

}

func (input *UpdateUserInput) MapUpdateInputToModel() {

}

func MapUserToDTO(user sqlc.User) *UserDTO {
	var age *int
	if user.UserAge != nil {
		v := int(*user.UserAge)
		age = &v
	}

	return &UserDTO{
		UUID:      user.UserUuid.String(),
		Name:      user.UserFullname,
		Email:     user.UserEmail,
		Age:       age,
		Status:    mapStatusText(int(user.UserStatus)),
		Level:     mapLevelText(int(user.UserLevel)),
		CreatedAt: user.UserCreatedAt.Format(time.DateTime),
		UpdatedAt: user.UserUpdatedAt.Format(time.DateTime),
	}
}

func MapUsersToDTO(users []sqlc.User) []UserDTO {
	dtos := make([]UserDTO, 0, len(users))

	for _, users := range users {
		dtos = append(dtos, *MapUserToDTO(users))
	}

	return dtos
}

func mapStatusText(status int) string {
	switch status {
	case 1:
		return "Show"
	case 2:
		return "Hide"
	default:
		return "None"
	}
}

func mapLevelText(status int) string {
	switch status {
	case 1:
		return "Admin"
	case 2:
		return "Member"
	default:
		return "None"
	}
}
