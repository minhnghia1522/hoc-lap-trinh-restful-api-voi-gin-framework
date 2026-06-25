package dto

import (
	"lesson08-prepare-connection/internal/db/sqlc"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id        int32     `json:"id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
}

func MapToResponseDto(model *sqlc.User) UserResponse {
	return UserResponse{
		Id:        model.UserID,
		Uuid:      model.Uuid,
		Name:      model.Name,
		Email:     model.Email,
		CreatedAt: model.CreatedAt.Format(time.DateTime),
	}
}
