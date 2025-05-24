package routes

import (
	"github.com/mobile-data-indonesia/inventaris-backend/handlers"
	"github.com/mobile-data-indonesia/inventaris-backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(r *gin.Engine) {
	itemService := services.NewItemService()
	itemHandler := handlers.NewUserHandler.NewItemHandler(itemService)

	itemGroup := r.Group("/items")
	{
		itemGroup.POST("/", itemHandler.CreateItem)
	}
}
