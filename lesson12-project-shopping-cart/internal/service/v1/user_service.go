package v1service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"user-management-api/internal/db/sqlc"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"
	"user-management-api/pkg/cache"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo  repository.UserRepository
	cache *cache.RedisCacheService
}

func NewUserService(repo repository.UserRepository, redisClient *redis.Client) UserService {
	return &userService{
		repo:  repo,
		cache: cache.NewRedisCacheService(redisClient),
	}
}

// CreateUser implements [UserService].
func (us *userService) CreateUser(ctx *gin.Context, userParam sqlc.CreateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()
	userParam.UserEmail = utils.NormalizeString(userParam.UserEmail)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userParam.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, utils.WrapError(err, "Failed to hash password", utils.ErrCodeInternal)
	}
	userParam.UserPassword = string(hashedPassword)
	userCreated, err := us.repo.CreateUser(context, userParam)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return sqlc.User{}, utils.NewError("email already existed", utils.ErrCodeConflict)
		}
		return sqlc.User{}, utils.WrapError(err, "Failed to insert user", utils.ErrCodeInternal)
	}
	if err := us.cache.Clear("users:*"); err != nil {
		log.Printf("Error clearing cache: %v", err)
	}
	return userCreated, nil
}

// DeleteUser implements [UserService].
func (us *userService) DeleteUser(ctx *gin.Context, uuid uuid.UUID) error {
	context := ctx.Request.Context()
	_, err := us.repo.TrashUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("User not found", utils.ErrCodeNotFound)
		}

		return utils.WrapError(err, "failed to restore user", utils.ErrCodeInternal)
	}
	if err := us.cache.Clear("users:*"); err != nil {
		log.Printf("Error clearing cache: %v", err)
	}
	return nil
}

// FindUserByUUID implements [UserService].
func (us *userService) FindUserByUUID(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()
	user, err := us.repo.GetUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found!", utils.ErrCodeNotFound)
		}
		return sqlc.User{}, utils.WrapError(err, "failed to get an user", utils.ErrCodeInternal)
	}
	return user, nil
}

// UpdateUser implements [UserService].
func (us *userService) UpdateUser(ctx *gin.Context, uuid uuid.UUID, updatedAt time.Time, params sqlc.UpdateUserParams) (sqlc.User, error) {
	context := ctx.Request.Context()
	var updatedUser sqlc.User

	err := us.repo.ExecTx(context, func(q *sqlc.Queries) error {
		var pgErr *pgconn.PgError
		user, err := q.GetUserForUpdateNoWait(context, uuid)
		if err != nil {
			if errors.As(err, &pgErr) && pgErr.Code == "55P03" {
				return utils.NewError("Data not available", utils.ErrCodeConflict)
			}
			return err
		}

		if !user.UserUpdatedAt.Equal(updatedAt) {
			return utils.NewError("User has updated before", utils.ErrCodeConflict)
		}
		time.Sleep(5 * time.Second)
		params.UserUuid = uuid
		updatedUser, err = q.UpdateUser(ctx, params)
		return err
	})

	if err != nil {
		return sqlc.User{}, err
	}
	if err := us.cache.Clear("users:*"); err != nil {
		log.Printf("Error clearing cache: %v", err)
	}
	return updatedUser, nil

}

// RestoreUser implements [UserService].
func (us *userService) RestoreUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()
	user, err := us.repo.RestoreUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
		}

		return sqlc.User{}, utils.WrapError(err, "failed to restore user", utils.ErrCodeInternal)
	}
	if err := us.cache.Clear("users:*"); err != nil {
		log.Printf("Error clearing cache: %v", err)
	}
	return user, nil
}

// SoftDeleteUser implements [UserService].
func (us *userService) SoftDeleteUser(ctx *gin.Context, uuid uuid.UUID) (sqlc.User, error) {
	context := ctx.Request.Context()
	user, err := us.repo.SoftDeleteUser(context, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
		}

		return sqlc.User{}, utils.WrapError(err, "failed to soft delete user", utils.ErrCodeInternal)
	}
	if err := us.cache.Clear("users:*"); err != nil {
		log.Printf("Error clearing cache: %v", err)
	}
	return user, nil
}

func (us *userService) GetAllUsers(ctx *gin.Context, search, orderBy, sort string, page, limit int32, deleted bool) ([]sqlc.User, int32, error) {
	context := ctx.Request.Context()

	if sort == "" {
		sort = "desc"
	}

	if orderBy == "" {
		orderBy = "user_created_at"
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limitInt := utils.GetIntEnv("LIMIT_ITEM_ON_PER_PAGE", 10)
		limit = int32(limitInt)
	}

	offset := (page - 1) * limit

	cacheKey := us.generateCacheKey(search, orderBy, sort, page, limit, deleted)
	var cachedData struct {
		Users []sqlc.User `json:"users"`
		Total int32       `json:"total"`
	}
	// Get data from cache if available
	if err := us.cache.Get(cacheKey, &cachedData); err == nil {
		log.Printf("Cache hit for key %s \n", cacheKey)
		return cachedData.Users, cachedData.Total, nil
	}

	users, err := us.repo.GetAllV2(context, search, orderBy, sort, limit, offset, deleted)

	if err != nil {
		return []sqlc.User{}, 0, utils.WrapError(err, "failed to fetch users", utils.ErrCodeInternal)
	}

	total, err := us.repo.CountUsers(context, sqlc.CountUsersParams{
		Search:  search,
		Deleted: &deleted,
	})
	if err != nil {
		return []sqlc.User{}, 0, utils.WrapError(err, "failed to count users", utils.ErrCodeInternal)
	}

	// Create cache data
	cacheData := struct {
		Users []sqlc.User `json:"users"`
		Total int32       `json:"total"`
	}{
		Users: users,
		Total: int32(total),
	}
	if err := us.cache.Set(cacheKey, cacheData, 5*time.Minute); err != nil {
		log.Printf("Failed to set cache for key %s: %v \n", cacheKey, err)
	}
	return users, int32(total), nil
}

func (us *userService) generateCacheKey(search, orderBy, sort string, page, limit int32, deleted bool) string {
	search = strings.TrimSpace(search)
	if search == "" {
		search = "none"
	}

	orderBy = strings.TrimSpace(orderBy)
	if orderBy == "" {
		orderBy = "user_created_at"
	}

	sort = strings.ToLower(strings.TrimSpace(sort))
	if sort == "" {
		sort = "desc"
	}

	return fmt.Sprintf("users:%s:%s:%s:%d:%d:%t", search, orderBy, sort, page, limit, deleted)
}
