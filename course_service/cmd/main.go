package main

import (
	pb "course/grpc/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Application struct {
	pb.UnimplementedCourseServiceServer
	Port string
}

func main() {
	app := &Application{
		Port: ":5004",
	}

	lis, err := net.Listen("tcp", app.Port)
	if err != nil {
		log.Fatalf("failed to listen on port 5004: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, app)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
