package store

import (
	"authentication/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type AuthStore struct {
	db *pgx.Conn
}

func (a *AuthStore) CreateUser(ctx context.Context, user models.User) bool {
	return false
}
