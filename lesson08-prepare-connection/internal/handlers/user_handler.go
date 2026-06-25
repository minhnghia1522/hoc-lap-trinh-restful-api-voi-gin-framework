package handlers

import (
	"lesson08-prepare-connection/internal/models"
	"lesson08-prepare-connection/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	uh.repo.FindById(id)
	ctx.JSON(http.StatusOK, gin.H{"data": "Get user by id"})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	uh.repo.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"data": "Created user"})
}
