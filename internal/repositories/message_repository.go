package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"time"

	"messaging/internal/model"
	"messaging/internal/repositories/base"
)

type MessageRepository struct {
	base.BaseSQLRepository[model.Message]
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{
		base.BaseSQLRepository[model.Message]{
			DB:        db,
			TableName: "messages",
			ScanFunc:  scanMessage,
		},
	}
}

func (r *MessageRepository) GetSentMessages() ([]model.Message, error) {
	rows, err := r.DB.Query("SELECT * FROM messages WHERE is_sent = true ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		msg, err := scanMessage(rows)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (r *MessageRepository) GetUnsentMessages(limit int) ([]model.Message, error) {
	rows, err := r.DB.Query(`
		SELECT id, to_phone, content FROM messages
		WHERE is_sent = false
		ORDER BY id ASC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.ID, &msg.ToPhone, &msg.Content); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (r *MessageRepository) MarkAsSent(id int64, sentAt time.Time, messageID uuid.UUID) error {
	_, err := r.DB.Exec(`
		UPDATE messages SET is_sent = true, sent_at = $1, message_id = $2 WHERE id = $3
	`, sentAt, messageID, id)
	return err
}

func scanMessage(rows *sql.Rows) (model.Message, error) {
	var msg model.Message
	err := rows.Scan(
		&msg.ID,
		&msg.Content,
		&msg.ToPhone,
		&msg.IsSent,
		&msg.CreatedAt,
		&msg.MessageID,
		&msg.SentAt,
	)
	return msg, err
}
