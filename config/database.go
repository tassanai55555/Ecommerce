package config

import (
    "fmt"
    "os"

    "github.com/tassanai55555/media-city-backend/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("❌ Failed to connect to database: " + err.Error())
    }     

    // Auto migrate your models
    db.AutoMigrate(&models.User{})

    DB = db
    fmt.Println("✅ Connected and Migrated Supabase DB")
}
