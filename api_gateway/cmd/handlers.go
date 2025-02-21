package main

import (
	"bytes"
	"context"
	"log"
	"net/http"

	video "gateway/grpc/client/video"
)

func (app *Application) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func (app *Application) UploadVideo(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	done := r.Header.Get("done")
	data := r.Body

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
		w.Write([]byte("Success"))
	} else {
		w.Write([]byte("Failed"))
	}
}

func (app *Application) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}
