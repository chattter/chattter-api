package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chattter/chattter-api/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello chattter!")

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
	}

	//================================================================================
	// Create the database connection
	//================================================================================

	// Get the datbase driver for the database string
	dbDriver := ParseDatabaseDriver(os.Getenv("DB"))
	if dbDriver == nil {
		log.Fatalln("Failed to create database driver. Check DB_URL environment variable")
	}

	// Create the database connection
	db, err := gorm.Open(dbDriver, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the database models
	db.AutoMigrate(
		&models.Account{},
	)

}
