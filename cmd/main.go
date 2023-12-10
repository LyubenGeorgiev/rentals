package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

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
		fmt.Println("Error converting port to integer:", err)
		return
	}

	// Database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := sql.Open("postgres", psqlInfo)
	for i := 0; err != nil && i < 5; i++ {
		time.Sleep(1 * time.Second)
		database, err = sql.Open("postgres", psqlInfo)
	}
	if err != nil {
		panic(err)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		panic(err)
	}
}
