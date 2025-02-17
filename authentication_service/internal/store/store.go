package store

import (
	"authentication/models"
	"context"
)

type Store interface {
	CreateUser(ctx context.Context, user models.User) bool
}

func NewStore() Store {

}
