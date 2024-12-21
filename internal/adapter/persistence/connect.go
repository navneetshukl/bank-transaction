package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connect to postgres
func ConnectToDB() {
	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database connection: %v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error connecting to the database: %v\n", err)
		return
	}
	fmt.Println("Successfully connected to PostgreSQL!")
}
