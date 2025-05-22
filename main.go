package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mobile-data-indonesia/inventaris-backend/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	config.ConnectDB()
 	
	port := os.Getenv("PORT")
	if err := router.Run(":" + port); err != nil {
		log.Println("Failed to run server")
	}
}