package config

import (
	"fmt"
	"log"
	"os"

	"Backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatalf("DATABASE_URL not set in environment")
    }

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Error getting DB from gorm DB: %v", err)
    }

    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

	if err := db.AutoMigrate(
		&models.Store{},
		&models.Category{},
		&models.Product{},
		&models.Address{},
		&models.Bank{},
		&models.User{},
	); err != nil {
		log.Fatalf("Error during other AutoMigrate: %v", err)
	}

    DB = db
    fmt.Println("Database connected successfully!")
}