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

func (r *RedisStore) DeleteEmailToken(ctx context.Context, token string) error {
	_, err := r.ds.Del(ctx, token).Result()
	if err != nil {
		log.Fatal("Failed to delete the email token")
		return err
	}

	return nil
}

func (r *RedisStore) GetEmailFromToken(ctx context.Context, token string) string {
	val := r.ds.Get(ctx, token).String()

	return val
}
