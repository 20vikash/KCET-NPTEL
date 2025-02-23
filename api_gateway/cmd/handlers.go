package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"time"

	video "gateway/grpc/client/video"
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

}
