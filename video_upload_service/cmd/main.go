package main

import (
	"log"
	"net/http"

	pb "video_upload/grpc/client/auth"
)

type Application struct {
	Port        string
	AuthService pb.AuthServiceClient
}

func main() {
	app := &Application{
		Port:        ":8080",
		AuthService: pb.ConnectToAuth(),
	}

	mux := app.handleRoutes()

	err := http.ListenAndServe(app.Port, mux)
	if err != nil {
		log.Println("Something went wrong in spinning up the server")
	}
}
