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
	web "gateway/web/components"
)

func (app *Application) Hello(w http.ResponseWriter, r *http.Request) {
	if app.Authorize.IsAuthenticated(r.Context()) {
		web.Layout(web.Home(app.Authorize.GetUserName(r.Context()))).Render(r.Context(), w)
	} else {
		web.Layout(web.Login()).Render(r.Context(), w)
	}
}

func (app *Application) Logout(w http.ResponseWriter, r *http.Request) {
	app.SessionManager.Destroy(r.Context())

	w.Header().Set("HX-Redirect", "/")
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	var user model.User

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.Bytes()

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
	}

	userDetails := &auth.UserDetails{
		UserName: user.UserName,
		Password: user.Password,
	}

	res, err := app.AuthService.LoginUser(ctx, userDetails)
	if err != nil {
		log.Println(err)
		web.Layout(web.Login()).Render(r.Context(), w)
		return
	}

	app.SessionManager.Put(r.Context(), "id", res.Id)
	app.SessionManager.Put(r.Context(), "user_name", res.UserName)
	app.SessionManager.Put(r.Context(), "role", res.Role)

	web.Layout(web.Home(app.Authorize.GetUserName(r.Context()))).Render(r.Context(), w)
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token := r.URL.Query().Get("token")

	_, err := app.AuthService.VerifyUser(ctx, &auth.Token{Token: token})
	if err != nil {
		w.Write([]byte("The Verification URL is expired"))
	} else {
		w.Write([]byte("Successfully verified your account."))
	}
}
