package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chattter/chattter-api/models"
	"github.com/chattter/chattter-api/services"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {

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

	// Setup all of the services
	socketsService := services.SocketsService{}

	// Create the server mux so we can handle both HTTP and WebSocket requests depending
	// on the request URL
	mux := http.NewServeMux()
	mux.Handle("/socket", &socketsService)
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world"))
	})

	// Create the HTTP server to host the API. By default, the server is bound to port 8080
	addr, ok := os.LookupEnv("ADDR")
	if !ok {
		addr = ":8080"
	}
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
