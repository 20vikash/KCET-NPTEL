package model

type User struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
