package handlers

import (
	"errors"
	"lesson08-prepare-connection/internal/models"
	"lesson08-prepare-connection/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid User ID",
		})
		return
	}
	user, err := uh.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
		"data": user,
	})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := uh.repo.Create(&user); err != nil {
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
		"data":    user,
	})
}
