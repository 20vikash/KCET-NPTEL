package store

import (
	"context"
	models "course/models/course"

	"github.com/jackc/pgx/v5"
)

type CourseStore struct {
	db *pgx.Conn
}

func (c *CourseStore) CreateCourse(ctx context.Context, course models.Course) bool {
	return true
}
