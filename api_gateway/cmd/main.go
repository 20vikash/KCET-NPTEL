package main

import (
	"log"
	"net/http"

	auth "gateway/grpc/client/auth"
	video "gateway/grpc/client/video"
)

type Application struct {
	Port         string
	AuthService  auth.AuthServiceClient
	VideoService video.VideoUploadServiceClient
}

func main() {
	app := &Application{
		Port:         ":8088",
		AuthService:  auth.ConnectToAuth(),
		VideoService: video.ConnectToVideo(),
	}

	mux := app.handleRoutes()

	err := http.ListenAndServe(app.Port, mux)
	if err != nil {
		log.Println(err)
	}
}
