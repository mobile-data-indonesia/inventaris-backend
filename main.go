package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mobile-data-indonesia/inventaris-backend/config"
	"github.com/mobile-data-indonesia/inventaris-backend/handlers"
	"github.com/mobile-data-indonesia/inventaris-backend/routes"
	"github.com/mobile-data-indonesia/inventaris-backend/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	userService := services.NewUserService(config.DB)
	userController := handlers.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userController)

	itemService := services.NewItemService(config.DB)
	itemController := handlers.NewItemHandler(itemService)
	itemRoutes := routes.NewItemRoutes(itemController)

	auditLogService := services.NewAuditLogService(config.DB)
	auditLogController := handlers.NewAuditLogHandler(auditLogService)
	auditLogRoutes := routes.NewAuditLogRoutes(auditLogController)

	router := gin.Default()

	userRoutes.RegisterRoutes(router)
	itemRoutes.RegisterRoutes(router)
	auditLogRoutes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if err := router.Run(":" + port); err != nil {
		log.Println("Failed to run server")
	}
}
