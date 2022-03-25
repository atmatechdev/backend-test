package database

import (
	"fmt"
	"log"
	"os"
	"technical-test-atmatech/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() (*gorm.DB, error) {
	// err := godotenv.Load(".env") // LOCAL DEVELOPMENT
	err := godotenv.Load("public.env") // FOR PRESENTATION / PUBLIC REPO
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// dsn := "host=localhost user=postgres password=-.-... dbname=atmatech-test post=5432 sslmode=disable Timezone=Asia/Bangkok"
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Postgres: ", err)
	}

	DB.AutoMigrate(&models.Book{}, &models.User{})

	return DB, err
}
