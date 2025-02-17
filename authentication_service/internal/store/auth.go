package store

import (
	"authentication/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthStore struct {
	db *pgx.Conn
}

func (a *AuthStore) CreateUser(ctx context.Context, user models.User) bool {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	sql := "INSERT INTO auth (email, user_name, password_hash) VALUES($1, $2, $3)"

	_, err := a.db.Exec(ctx, sql, user.Email, user.UserName, password)
	if err != nil {
		log.Panic(err)
		return false
	}

	return true
}
