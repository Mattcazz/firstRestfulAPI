package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Env file sucks!")
	}

	dbHost := os.Getenv("HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")

	fmt.Printf("%v", dbUser)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Not correctly connected")
	} else {
		fmt.Println("Connected to the DataBase called", dbName)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to close the database connection:", err)
	} else {
		fmt.Println("Connection closed successfully")
	}
	sqlDB.Close()
}
