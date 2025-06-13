CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    to_phone VARCHAR(20) NOT NULL,
    content TEXT NOT NULL,
    is_sent BOOLEAN DEFAULT FALSE,
    sent_at TIMESTAMP,
    message_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
