package repositories

import (
	"database/sql"

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
