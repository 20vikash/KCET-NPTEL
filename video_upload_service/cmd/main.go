package main

import (
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
}
