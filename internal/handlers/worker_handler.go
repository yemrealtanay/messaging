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

func (h *WorkerHandler) Start(w http.ResponseWriter, r *http.Request) {
	h.Worker.Start(h.Service.SendUnsentMessages)
	response.JSON(w, http.StatusOK, "Worker has been started", struct{}{})
}

func (h *WorkerHandler) Stop(w http.ResponseWriter, r *http.Request) {
	h.Worker.Stop()
	response.JSON(w, http.StatusOK, "Worker has been stopped", struct{}{})
}
