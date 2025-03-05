package model

type User struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserSession struct {
	UserName string
	Id       string
	Role     string
}
