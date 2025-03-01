package video

import (
	"log"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToVideo() VideoUploadServiceClient {
	conn, err := grpc.NewClient("video_upload_service:5002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("Failed to connect to auth service")
	}

	c := NewVideoUploadServiceClient(conn)

	return c
}
