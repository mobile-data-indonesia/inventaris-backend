package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mobile-data-indonesia/inventaris-backend/handlers"
)

type UserRoutes struct {
	UserController *handlers.UserHandler
}

func NewUserRoutes(ctrl *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{UserController: ctrl}
}

func (r *UserRoutes) RegisterRoutes(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.POST("/register", r.UserController.Register)
		user.POST("/login", r.UserController.Login)
		user.PUT("/:id", r.UserController.UpdateUser)
		user.GET("/:id", r.UserController.GetUserByID)
		user.GET("/", r.UserController.GetAllUsers)
		user.POST("/refresh-token", r.UserController.RefreshToken)
	}
}
