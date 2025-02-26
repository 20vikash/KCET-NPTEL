package auth

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToAuth() AuthServiceClient {
	conn, err := grpc.NewClient("auth_service:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("Failed to connect to auth service")
	}

	c := NewAuthServiceClient(conn)

	return c
}
