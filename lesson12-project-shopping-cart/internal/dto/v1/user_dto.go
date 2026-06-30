package v1dto

import (
	"time"
	"user-management-api/internal/db/sqlc"
	"user-management-api/internal/utils"
)

type UserDTO struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"full_name"`
	Email     string    `json:"email_address"`
	Age       *int      `json:"age,omitempty"`
	Status    string    `json:"status"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *string   `json:"deleted_at,omitempty"`
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"omitempty,gt=1,lt=150"`
	Password string `json:"password" binding:"required,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UpdateUserInput struct {
	Name      string    `json:"name" binding:"required"`
	Age       int       `json:"age" binding:"omitempty,gt=1,lt=150"`
	Status    int       `json:"status" binding:"required,oneof=1 2"`
	Level     int       `json:"level" binding:"required,oneof=1 2"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   int32  `form:"page" binding:"omitempty,gte=1"`
	Limit  int32  `form:"limit" binding:"omitempty,gte=1,lte=500"`
	Order  string `form:"order_by" binding:"omitempty,oneof=user_id user_created_at"`
	Sort   string `form:"sort" binding:"omitempty,oneof=asc desc"`
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

func (input *UpdateUserInput) MapUpdateInputToModel() sqlc.UpdateUserParams {
	return sqlc.UpdateUserParams{
		UserFullname: &input.Name,
		UserAge:      utils.ConvertToInt32Pointer(input.Age),
		UserStatus:   utils.ConvertToInt32Pointer(input.Status),
		UserLevel:    utils.ConvertToInt32Pointer(input.Level),
	}
}

func MapUserToDTO(user sqlc.User) *UserDTO {
	var age *int
	if user.UserAge != nil {
		v := int(*user.UserAge)
		age = &v
	}

	var deletedAt *string
	if user.UserDeletedAt != nil {
		v := user.UserDeletedAt.Format(time.DateTime)
		deletedAt = &v
	}

	return &UserDTO{
		UUID:      user.UserUuid.String(),
		Name:      user.UserFullname,
		Email:     user.UserEmail,
		Age:       age,
		Status:    mapStatusText(int(user.UserStatus)),
		Level:     mapLevelText(int(user.UserLevel)),
		CreatedAt: user.UserCreatedAt,
		UpdatedAt: user.UserUpdatedAt,
		DeletedAt: deletedAt,
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
		return "Active"
	case 2:
		return "Inactive"
	case 3:
		return "Banned"
	default:
		return "None"
	}
}

func mapLevelText(status int) string {
	switch status {
	case 1:
		return "Administrator"
	case 2:
		return "Moderator"
	case 3:
		return "Member"
	default:
		return "None"
	}
}
