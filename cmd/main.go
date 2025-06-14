package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"messaging/internal/handlers"
	"messaging/internal/repositories"
	"messaging/internal/router"
	"messaging/internal/services"
	"messaging/internal/worker"
	"net/http"
	"os"

	"context"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"messaging/internal/db"
)

func main() {
	_ = godotenv.Load()
	ctx := context.Background()

	conn, err := db.NewConnection()

	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	log.Println("connected to postgres successfully")
	defer conn.Close()

	repo := repositories.NewMessageRepository(conn)
	service := services.NewMessageService(repo)
	w := worker.NewWorker()
	w.Start(service.SendUnsentMessages)

	reg := &router.HandlerRegistry{
		Message: handlers.NewMessageHandler(repo),
		Worker:  handlers.NewWorkerHandler(w, service),
	}
	r := router.NewRouter(reg)

	//redis

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	log.Println("connected to redis successfully")

	log.Println("listening on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
