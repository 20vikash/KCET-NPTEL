package main

import (
	"log"
	"net"
	processing "video_upload/grpc/client"
	pb "video_upload/grpc/server"

	"google.golang.org/grpc"
)

type Application struct {
	pb.UnimplementedVideoUploadServiceServer
	Port         string
	VideoProcess processing.VideoProcessingServiceClient
}

func main() {
	app := &Application{
		Port:         ":5002",
		VideoProcess: processing.ConnectToVideoProcessingService(),
	}

	lis, err := net.Listen("tcp", app.Port)
	if err != nil {
		log.Fatalf("failed to listen on port 5002: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVideoUploadServiceServer(s, app)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
