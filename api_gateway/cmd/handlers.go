package main

import "net/http"

func (app *Application) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func (app *Application) UploadVideo(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) CreateUser(w http.ResponseWriter, r *http.Request) {

}
