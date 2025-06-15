package handlers

import (
	"net/http"

	"messaging/internal/response"
	"messaging/internal/services"
	"messaging/internal/worker"
)

type WorkerHandler struct {
	Worker  *worker.Worker
	Service *services.MessageService
}

func NewWorkerHandler(w *worker.Worker, s *services.MessageService) *WorkerHandler {
	return &WorkerHandler{
		Worker:  w,
		Service: s,
	}
}

// Start godoc
// @Summary Starts the background worker
// @Description Triggers the background worker to begin processing unsent messages
// @Tags worker
// @Produce json
// @Success 200 {object} response.EmptyResponse
// @Router /start [post]
func (h *WorkerHandler) Start(w http.ResponseWriter, r *http.Request) {
	h.Worker.Start(h.Service.SendUnsentMessages)
	response.JSON(w, http.StatusOK, "Worker has been started", struct{}{})
}

// Stop godoc
// @Summary Stops the background worker
// @Description Gracefully stops the background message processing worker
// @Tags worker
// @Produce json
// @Success 200 {object} response.EmptyResponse
// @Router /stop [post]
func (h *WorkerHandler) Stop(w http.ResponseWriter, r *http.Request) {
	h.Worker.Stop()
	response.JSON(w, http.StatusOK, "Worker has been stopped", struct{}{})
}
