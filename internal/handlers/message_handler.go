package handlers

import (
	"encoding/json"
	"messaging/internal/repositories"
	"net/http"
)

type MessageHandler struct {
	Repo *repositories.MessageRepository
}

func NewMessageHandler(repo *repositories.MessageRepository) *MessageHandler {
	return &MessageHandler{Repo: repo}
}

func (h *MessageHandler) GetAllSentMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Repo.GetSentMessages()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		println("GetSentMessages error:", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
