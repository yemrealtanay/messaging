package app

import (
	"fmt"
	"log"
	"messaging/internal/config"
	"messaging/internal/db"
	"messaging/internal/handlers"
	"messaging/internal/logs"
	"messaging/internal/repositories"
	"messaging/internal/router"
	"messaging/internal/services"
	"messaging/internal/worker"
	"net/http"
)

func Run() error {

	// Config bölümü
	//aslında bu kadar küçük bir task için biraz fazla görünebilir.
	//yine de env'leri karmaşık bir şekilde çağırmak gözümü rahatsız ediyor.
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config load error: %w", err)
	}

	// DB bağlantısı
	conn, err := db.NewPgConnection(cfg)
	if err != nil {
		return fmt.Errorf("postgres error: %w", err)
	}
	defer conn.Close()
	log.Println("Postgres connected")

	// Redis
	rdb, err := db.NewRedisClient(cfg)
	if err != nil {
		return fmt.Errorf("redis init error: %w", err)
	}
	log.Println("Redis connected")

	// Logger, repository, service
	logger := logs.NewRedisLogger(rdb, "message_logs")
	repo := repositories.NewMessageRepository(conn)
	service := services.NewMessageService(repo, logger, cfg.WebhookURL)
	worker := worker.NewWorker()
	worker.Start(service.SendUnsentMessages)

	// Router
	reg := &router.HandlerRegistry{
		Message: handlers.NewMessageHandler(repo),
		Worker:  handlers.NewWorkerHandler(worker, service),
	}
	r := router.NewRouter(reg)

	log.Println("listening on :8080")
	return http.ListenAndServe(":8080", r)
}
