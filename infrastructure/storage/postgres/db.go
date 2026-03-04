package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitConnection() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read environment variables or provide default values
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the PostgreSQL connection string
	postgresConnectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, user, password, dbname, "Asia/Tashkent")

	// Open a connection to the database
	var err error
	db, err = sql.Open("postgres", postgresConnectionString)
	if err != nil {
		return err
	}

	// Set the maximum number of open connections
	db.SetMaxOpenConns(1024)

	// Set the maximum number of idle connections
	db.SetMaxIdleConns(128)

	// Set the connection expiration time
	db.SetConnMaxLifetime(5 * time.Minute)

	// Ping the database to check if the connection is successful
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

func CloseConnection() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func QueryRow(query string, params ...any) *sql.Row {
	//defer CloseConnection()
	return db.QueryRow(query, params...)
}

func Query(query string, params ...any) (*sql.Rows, error) {
	//defer CloseConnection()
	return db.Query(query, params...)
}

func Exec(query string, params ...any) (sql.Result, error) {

	return db.Exec(query, params...)
}
