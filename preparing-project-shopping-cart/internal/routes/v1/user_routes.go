package v1routes

import (
	"user-management-api/internal/handler/v1"

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
		route.POST("", ur.handler.CreateUser)
		route.PUT("/:uuid", ur.handler.UpdateUser)
		route.DELETE("/:uuid", ur.handler.DeleteUser)
	}
}
