package store

import (
	"context"
	models "course/models/course"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	Course interface {
		CreateCourse(ctx context.Context, course models.Course) bool
	}
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{
		Course: &CourseStore{db: db},
	}
}
