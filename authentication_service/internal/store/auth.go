package store

import (
	"authentication/models"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgconn"
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				fmt.Println("Error: Duplicate entry (unique constraint violation)")
			}
		}

		return false
	}

	return true
}

func (a *AuthStore) VerifyUser(ctx context.Context, email string) error {
	sql := "UPDATE auth SET is_activated=true WHERE email=$1"

	_, err := a.db.Exec(ctx, sql, email)
	if err != nil {
		fmt.Printf("Cannot verify the user with the email %s", email)
		return err
	}

	return nil
}

func (a *AuthStore) LoginUser(ctx context.Context, user string, password string) (bool, error) {
	sql := "SELECT is_activated FROM auth WHERE user=$1"
	var is_activated bool

	err := a.db.QueryRow(ctx, sql, user).Scan(&is_activated)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if !is_activated {
		return false, errors.New("verify")
	}

	return true, nil
}
