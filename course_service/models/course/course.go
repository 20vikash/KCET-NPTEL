package models

type Course struct {
	Title       string
	Description string
	TeacherId   int
}

type Enrollment struct {
	UserId   int
	CourseId int
}
