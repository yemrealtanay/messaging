package main

import (
	"log"
	"os"
	"strconv"

	"messaging/internal/config"
	"messaging/internal/db"
	"messaging/internal/faker"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %w", err)
	}

	count := 5000
	if len(os.Args) > 1 {
		if parsed, err := strconv.Atoi(os.Args[1]); err == nil {
			count = parsed
		}
	}

	conn, err := db.NewPgConnection(cfg)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer conn.Close()

	log.Printf("⚙Inserting %d fake messages...", count)
	if err := faker.GenerateFakeMessages(conn, count); err != nil {
		log.Fatalf("Faker failed: %v", err)
	}

	log.Println("Faker inserted messages successfully.")
}
