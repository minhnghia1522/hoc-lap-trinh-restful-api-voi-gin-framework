package v1routes

import (
	v1handler "user-management-api/internal/handler/v1"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *v1handler.UserHandler
}

func NewUserRoutes(handler *v1handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	route := r.Group("/users")
	{
		route.GET("", ur.handler.Search)
		route.GET("/:uuid", ur.handler.GetUserByUUID)
		route.POST("", ur.handler.PostCreateUser)
		route.PUT("/:uuid", ur.handler.PutUpdateUser)
		route.PUT("/:uuid/restore", ur.handler.RestoreUser)
		route.DELETE("/:uuid", ur.handler.SoftDeleteUser)
		route.DELETE("/:uuid/trash", ur.handler.DeleteUser)
		route.GET("/me", ur.handler.GetMe)
	}
}
