package store

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	ds *redis.Client
}

func (r *RedisStore) SetEmailToken(ctx context.Context, email string, token string) error {
	err := r.ds.Set(ctx, token, "email:"+email, 24*time.Hour).Err()
	if err != nil {
		log.Fatal("Failed to set email token")
		return err
	}

	return nil
}
