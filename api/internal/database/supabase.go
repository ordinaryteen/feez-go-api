package database

import (
	"context"
	"fmt" // Kita butuh 'fmt' untuk ngebangun string
	"log"
	"net/url" // Kita butuh 'url' untuk ngamanin password
	"os"

	"github.com/jackc/pgx/v5/pgxpool" // <-- Kita TETEP pake 'pgxpool', bukan 'pgx'
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found.")
	}

	dbUser := os.Getenv("DB_USER")
	rawPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || rawPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("One or more database environment variables are not set")
	}

	encodedPassword := url.QueryEscape(rawPassword)

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		dbUser,
		encodedPassword,
		dbHost,
		dbPort,
		dbName,
	)

	conn, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Database PING failed:", err)
	}

	DB = conn
	log.Println("Database connection established and PINGED successfully!")
}
