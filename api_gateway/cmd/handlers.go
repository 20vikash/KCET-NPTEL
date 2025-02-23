package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	auth "gateway/grpc/client/auth"
	video "gateway/grpc/client/video"
	model "gateway/models/auth"
)

func (app *Application) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func (app *Application) UploadVideo(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Ok")

	done := r.Header.Get("done")
	data := r.Body

	defer r.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	dataBytes := buf.Bytes()

	videoData := &video.VideoData{
		Data: dataBytes,
		Done: done,
	}

	_, err := app.VideoService.UploadBinary(ctx, videoData)
	if err != nil {
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (app *Application) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user model.User

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.Bytes()

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
	}

	userDetails := &auth.UserDetails{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
	}

	_, err := app.AuthService.CreateUser(ctx, userDetails)
	if err != nil {
		log.Println(err)
	}
}

func (app *Application) VerifyUser(w http.ResponseWriter, r *http.Request) {

}
