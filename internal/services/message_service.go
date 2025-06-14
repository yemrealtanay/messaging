package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"

	"messaging/internal/model"
	"messaging/internal/repositories"
)

type MessageService struct {
	Repo *repositories.MessageRepository
}

func NewMessageService(repo *repositories.MessageRepository) *MessageService {
	return &MessageService{Repo: repo}
}

func (s *MessageService) SendUnsentMessages() error {
	messages, err := s.Repo.GetUnsentMessages(2)
	if err != nil {
		return fmt.Errorf("db query is failed: %w", err)
	}

	for _, msg := range messages {
		if err := s.sendMessage(msg); err != nil {
			return fmt.Errorf("send message failed: %w", err)
			continue
		}
	}

	return nil
}

func (s *MessageService) sendMessage(msg model.Message) error {
	payload := map[string]string{
		"to":      msg.ToPhone,
		"content": msg.Content,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "https://webhook.site/b31e1d64-f0e9-4dd8-8505-69d29a5df172", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send message failed: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("send message failed: %s", resp.Status)
	}

	err = s.Repo.MarkAsSent(msg.ID, time.Now(), uuid.New())

	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}

	log.Printf("✅ sent message ID %d\n", msg.ID)
	return nil
}
