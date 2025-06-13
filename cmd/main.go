package main

import (
	"fmt"
	"log"
	handler "messaging/internal/handler/message"
	"messaging/internal/repository/message"
	"messaging/internal/router"
	"net/http"
	"os"

	"context"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"messaging/internal/db"
)

func main() {
	ctx := context.Background()

	conn, err := db.NewConnection()

	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	log.Println("connected to postgres successfully")
	defer conn.Close()

	messageRepo := message.NewMessageRepository(conn)
	messageHandler := handler.NewHandler(messageRepo)

	r := router.NewRouter(messageHandler)
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
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
