package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	pgURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", pgURL)
		if err != nil {
			log.Printf("sql.Open failed: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		log.Printf("attempt %d: failed to connect to postgres: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("unable to connect to postgres after retries: %w", err)
}
