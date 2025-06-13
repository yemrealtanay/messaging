package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	//pgsql
	pgURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	log.Println("connected to postgres successfully")

	//redis

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	log.Println("connected to redis successfully")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Fatalf("failed to write response: %v", err)
		}
	})

	log.Println("listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
