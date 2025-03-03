package models

type User struct {
	Id          int
	Email       string
	UserName    string
	Password    string
	IsActivated bool
	Role        string
}
