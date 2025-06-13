package message

import (
	"encoding/json"
	"messaging/internal/repository/message"
	"net/http"
)

type Handler struct {
	Repo *message.Repository
}

func NewHandler(repo *message.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) GetAllSentMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Repo.GetSentMessages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
