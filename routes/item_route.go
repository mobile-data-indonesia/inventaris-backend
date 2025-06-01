package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mobile-data-indonesia/inventaris-backend/handlers"
)

type ItemRoutes struct {
	ItemController *handlers.ItemHandler
}

func NewItemRoutes(ctrl *handlers.ItemHandler) *ItemRoutes {
	return &ItemRoutes{ItemController: ctrl}
}

func (r *ItemRoutes)RegisterRoutes(router *gin.Engine) {
	item := router.Group("/items")
	{
		item.POST("/", r.ItemController.CreateItem)
		item.PUT("/:id", r.ItemController.UpdateItem)
		item.GET("/:id", r.ItemController.GetItemByID)
		item.GET("/", r.ItemController.GetAllItems)
	}
}
