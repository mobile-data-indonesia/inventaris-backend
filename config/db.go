package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/mobile-data-indonesia/inventaris-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func ConnectDB() {
	once.Do(func() {
		dsn := os.Getenv("DB_CONFIG")
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Gagal koneksi ke database: %v", err))
		}
		fmt.Println("Berhasil koneksi ke database")

		DB.AutoMigrate(&models.User{})
		DB.AutoMigrate(&models.Item{})
		fmt.Println("Database Migrated")
	})
}
