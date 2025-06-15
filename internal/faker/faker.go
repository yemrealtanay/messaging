package faker

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

var SampleContents = []string{
	"Your verification code is 1234.",
	"Welcome to the platform!",
	"Please reset your password.",
	"Your delivery is on the way.",
	"Thanks for signing up!",
	"Reminder: your appointment is tomorrow.",
	"Discount code: SAVE20",
}

func GenerateFakeMessages(db *sql.DB, count int) error {
	stmt, err := db.Prepare(`
		INSERT INTO messages (to_phone, content, is_sent, created_at)
		VALUES ($1, $2, false, $3)
	`)

	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()
	for i := 0; i < count; i++ {
		phone := fmt.Sprintf("+90555%06d", rand.Intn(1000000))
		content := SampleContents[rand.Intn(len(SampleContents))]
		createdAt := time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour)

		if _, err := stmt.Exec(phone, content, createdAt); err != nil {
			return fmt.Errorf("insert message: %w", err)
		}
	}
	return nil
}
