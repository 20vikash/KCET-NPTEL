package store

import (
	"authentication/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	Auth interface {
		CreateUser(context.Context, models.User) bool
	}

	Redis interface {
		SetEmailToken(context.Context)
	}
}

func NewStore(db *pgx.Conn, redis *redis.Client) *Store {
	return &Store{
		Auth:  &AuthStore{db},
		Redis: &RedisStore{redis},
	}
}
