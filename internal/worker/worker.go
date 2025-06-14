package worker

import (
	"context"
	"log"
	"time"
)

type Worker struct {
	ctx     context.Context
	cancel  context.CancelFunc
	ticker  *time.Ticker
	running bool
}

func NewWorker() *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		ctx:    ctx,
		cancel: cancel,
		ticker: time.NewTicker(30 * time.Second),
	}
}

func (w *Worker) Start(task func() error) {
	if w.running {
		log.Println("Worker already running")
		return
	}

	log.Println("🚀 Starting worker...")

	ctx, cancel := context.WithCancel(context.Background())
	w.ctx = ctx
	w.cancel = cancel
	w.ticker = time.NewTicker(30 * time.Second)
	w.running = true

	go func() {
		log.Println("Worker is running")
		for {
			select {
			case <-w.ctx.Done():
				log.Println("Worker is shutting down")
				w.ticker.Stop()
				w.running = false
				return
			case <-w.ticker.C:
				log.Println("processing work")
				if err := task(); err != nil {
					log.Printf("error processing: %v\n", err)
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	if !w.running {
		log.Println("Worker is not running")
		return
	}

	log.Println("Stopping worker...")
	w.cancel()
}

func (w *Worker) IsRunning() bool {
	return w.running
}
