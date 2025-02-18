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
		VerifyUser(ctx context.Context, email string) error
	}

	Redis interface {
		SetEmailToken(ctx context.Context, email string, token string) error
		GetEmailFromToken(ctx context.Context, token string) string
	}
}

func NewStore(db *pgx.Conn, redis *redis.Client) *Store {
	return &Store{
		Auth:  &AuthStore{db},
		Redis: &RedisStore{redis},
	}
}
