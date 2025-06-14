package handlers

import (
	"messaging/internal/repositories"
	"messaging/internal/response"
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
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(messages) == 0 {
		response.JSON(w, http.StatusOK, "No Message Data", struct{}{})
	} else {
		response.JSON(w, http.StatusOK, "All Sent Messages retrieved successfully", messages)
	}
}
