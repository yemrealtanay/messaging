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

// GetAllSentMessages godoc
// @Summary Returns all sent messages
// @Description Fetches messages where is_sent=true
// @Tags messages
// @Produce json
// @Success 200 {object} response.MessageListResponse
// @Failure 500 {object} response.EmptyResponse
// @Router /messages [get]
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
