package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Truong62/taskoria/auth-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	DB = db
	fmt.Println("✅ Connected to DB successfully")
	return db
}
