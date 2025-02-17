package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"video_upload/grpc/client/auth"
)

func (app *Application) rootPage(w http.ResponseWriter, r *http.Request) {
	done := r.Header.Get("done")
	isDone, _ := strconv.ParseBool(done)

	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, "Desktop", "videos", "image.mp4")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Cannot open the file")
	}
	defer file.Close()
	body := r.Body

	if isDone {
		ConvertToHls(filePath)
		log.Println("Done is true")
	}

	fmt.Println(body)
	_, err = io.Copy(file, body)
	if err != nil {
		log.Println("Error. yea.. pretty much")
	}

	w.Write([]byte("Success"))
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

func (app *Application) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Something went wrong")
	}
	defer r.Body.Close()

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal("Something went wrong")
	}

	user_ := &auth.UserDetails{
		Email:    user.Email,
		Password: user.Password,
		UserName: user.UserName,
	}

	_, err = app.AuthService.CreateUser(ctx, user_)
	if err != nil {
		log.Panic("Cannot create an user")
	} else {
		log.Println("Created an user")
	}
}
