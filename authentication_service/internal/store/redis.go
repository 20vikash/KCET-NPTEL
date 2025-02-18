package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	ds *redis.Client
}

func (r *RedisStore) SetEmailToken(ctx context.Context) {

}
