package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type PG struct {
	Host     string
	Username string
	Password string
	Database string
}

func (p *PG) Connect(ctx context.Context) *pgx.Conn {
	pgUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", p.Username, p.Password, p.Host, p.Database)

	conn, err := pgx.Connect(ctx, pgUrl)
	if err != nil {
		log.Panic("Cannot connect to the database")
	}

	return conn
}

func (p *PG) CreateUser(ctx context.Context) bool {
	return false
}
