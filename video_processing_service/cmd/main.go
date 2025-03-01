package main

import (
	"log"
	"net"
	pb "video_processing/grpc/server"

	"google.golang.org/grpc"
)

type Application struct {
	pb.UnimplementedVideoProcessingServiceServer
	Port string
}

func main() {
	app := &Application{
		Port: ":5003",
	}

	lis, err := net.Listen("tcp", app.Port)
	if err != nil {
		log.Fatalf("failed to listen on port 5003: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVideoProcessingServiceServer(s, app)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
