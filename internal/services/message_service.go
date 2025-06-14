package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"log"
	"net/http"
	"os"
	"time"

	"messaging/internal/logs"
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

	operation := func() error {
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

		//burada bir backoff stratejisi göstermek istedim. Geçici http hatalarında tekrar denemek
		//verinin işlendiğinden emin olmamızı sağlayabilir.
		//test etmek için env'den webhook url değiştirebilirsiniz.
		//Gönderim başarısız olursa mesaj 30 saniye boyunca exponential delay ile tekrar denenir.
		//404 ve 400 gibi fatal hatalarda tekrar deneme yapılmaz.
		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest {
				return backoff.Permanent(fmt.Errorf("fatal status: %s", resp.Status))
			}
			return fmt.Errorf("retryable status: %s", resp.Status)
		}

		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			return fmt.Errorf("failed to decode webhook response: %w", err)
		}

		return nil
	}

	expBackOff := backoff.NewExponentialBackOff()
	expBackOff.MaxElapsedTime = 30 * time.Second

	if err := backoff.Retry(operation, expBackOff); err != nil {
		return fmt.Errorf("send message failed after retries: %w", err)
	}

	if err := s.Repo.MarkAsSent(msg.ID, time.Now(), respBody.MessageID); err != nil {
		return fmt.Errorf("mark message as sent failed after retries: %w", err)
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
