package handlers

import (
	"database/sql"
	"errors"
	"lesson08-prepare-connection/dto"
	"lesson08-prepare-connection/internal/db/sqlc"
	"lesson08-prepare-connection/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var err error
	uuidParam := ctx.Param("uuid")

	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid UUID",
		})
		return
	}
	user, err := uh.repo.FindByUUID(ctx, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "A system error has occurred.",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.MapToResponseDto(&user),
	})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var userParam sqlc.CreateUserParams
	if err := ctx.ShouldBindBodyWithJSON(&userParam); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := uh.repo.Create(ctx, userParam)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": "Email already exist",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Created user",
		"data":    dto.MapToResponseDto(&user),
	})
}
