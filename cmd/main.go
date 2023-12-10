package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/LyubenGeorgiev/rentals/db"
	"github.com/LyubenGeorgiev/rentals/internal"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Fetch environment variables
	host := os.Getenv("DATABASE_HOST")
	portStr := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Convert port string to integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Error converting port to integer:", err)
	}

	// Database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := sql.Open("postgres", psqlInfo)
	// Wait for postgres to boot up completely
	for i := 0; err != nil && i < 5; i++ {
		time.Sleep(1 * time.Second)
		database, err = sql.Open("postgres", psqlInfo)
	}
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer database.Close()

	// Ping the database
	err = database.Ping()
	if err != nil {
		log.Fatal("Could not ping the database:", err)
	}

	// Create postgres storage
	postgresStorage := db.NewPostgresStorage(database)

	// Create a new RentalsHandler
	handler := internal.NewRentalsHandler(postgresStorage)

	// Create a new router using gorilla/mux
	router := mux.NewRouter()

	// Handle routes with path variables
	router.HandleFunc("/rentals/{id}", handler.GetRentalByID).Methods("GET")

	// Handle route for /rentals with optional query parameters
	router.HandleFunc("/rentals", handler.GetRentalByQuery).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port 8080...")
	if err = server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
