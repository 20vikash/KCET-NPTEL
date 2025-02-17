package store

import (
	"authentication/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	Auth interface {
		CreateUser(context.Context, models.User) bool
	}
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{
		Auth: &AuthStore{db},
	}
}
