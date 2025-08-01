package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Truong62/taskoria/auth-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	createDatabaseIfNotExists()

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

func createDatabaseIfNotExists() {
	// Connect to Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER_AUTH"),
		os.Getenv("DB_PASSWORD_AUTH"),
		os.Getenv("DB_PORT_AUTH"),
	)

	drivenName := "postgres"
	db, err := sql.Open(drivenName, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL server: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	dbName := os.Getenv("DB_NAME_AUTH")

	// Check if the database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1)", dbName).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}

	if exists {
		log.Printf("Database already exists: %s", dbName)
		return
	}

	// Create database if not found
	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Printf("Error creating database: %v", err)
	} else {
		log.Println("Database created:", dbName)
	}
}
