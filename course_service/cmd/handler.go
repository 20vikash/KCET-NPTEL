package main

import (
	"context"
	course "course/grpc/server"
)

func (app *Application) CreateCourse(context.Context, *course.CourseData) (*course.CourseResponse, error) {
	return &course.CourseResponse{
		Message: "Hello",
	}, nil
}
