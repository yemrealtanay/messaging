package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisLogger struct {
	Client *redis.Client
	Key    string
}

func NewRedisLogger(client *redis.Client, key string) *RedisLogger {
	return &RedisLogger{
		Client: client,
		Key:    key,
	}
}

func (l *RedisLogger) LogMessage(ctx context.Context, msg any) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return l.Client.LPush(ctx, l.Key, string(payload)).Err()
}
