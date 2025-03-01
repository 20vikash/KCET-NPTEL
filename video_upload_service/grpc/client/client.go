package processing

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToVideoProcessingService() VideoProcessingServiceClient {
	conn, err := grpc.NewClient("video_process:5003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("Failed to connect to video processing service")
	}

	c := NewVideoProcessingServiceClient(conn)

	return c
}
