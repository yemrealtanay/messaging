package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"messaging/internal/logs"
	"net/http"
	"os"
	"time"

	"messaging/internal/model"
	"messaging/internal/repositories"
)

type MessageService struct {
	Repo   *repositories.MessageRepository
	Logger *logs.RedisLogger
}

func NewMessageService(repo *repositories.MessageRepository, l *logs.RedisLogger) *MessageService {
	return &MessageService{
		Repo:   repo,
		Logger: l,
	}
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
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		return fmt.Errorf("WEBHOOK_URL env not set")
	}

	var respBody struct {
		Message   string `json:"message"`
		MessageID string `json:"messageId"`
	}

	payload := map[string]string{
		"to":      msg.ToPhone,
		"content": msg.Content,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(body))
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

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return fmt.Errorf("failed to decode webhook response: %w", err)
	}

	err = s.Repo.MarkAsSent(msg.ID, time.Now(), respBody.MessageID)

	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}

	logPayload := map[string]any{
		"messageId": respBody.MessageID,
		"to":        msg.ToPhone,
		"content":   msg.Content,
		"sent_at":   time.Now(),
	}

	if err := s.Logger.LogMessage(context.Background(), logPayload); err != nil {
		log.Printf("failed to log message: %s", err)
	}

	log.Printf("sent message ID %d\n", msg.ID)
	return nil
}
