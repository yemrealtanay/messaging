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

func (w *Worker) Start(process func() error) {
	if w.running {
		log.Println("worker is already running")
		return
	}
	w.running = true

	go func() {
		log.Println("worker is running")
		for {
			select {
			case <-w.ctx.Done():
				log.Println("worker is shutting down")
				return
			case <-w.ticker.C:
				log.Println("processing work")
				if err := process(); err != nil {
					log.Println("processing error:", err)
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	if !w.running {
		log.Println("worker is already stopped")
		return
	}
	w.ticker.Stop()
	w.cancel()
	w.running = false
}
