package store

import (
	"context"
	models "course/models/course"
	"log"

	"github.com/jackc/pgx/v5"
)

type CourseStore struct {
	db *pgx.Conn
}

func (c *CourseStore) CreateCourse(ctx context.Context, course models.Course) error {
	sql := "INSERT INTO course (title, description, thumbnail) VALUES ($1, $2, $3) RETURNING id"
	var id int64

	err := c.db.QueryRow(ctx, sql, course.Title, course.Description, "t").Scan(&id)
	if err != nil {
		log.Println(err)
		return err
	}

	sql = "INSERT INTO course_teacher (course_id, teacher_id) VALUES ($1, $2)"

	_, err = c.db.Exec(ctx, sql, id, course.TeacherId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *CourseStore) EnrollStudent(ctx context.Context, enrollment models.Enrollment) error {
	sql := "INSERT INTO enrollment (user_id, course_id) VALUES ($1, $2)"

	_, err := c.db.Exec(ctx, sql, enrollment.UserId, enrollment.CourseId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
