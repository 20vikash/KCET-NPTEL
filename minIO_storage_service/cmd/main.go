package main

import (
	"log"
	"net"

	pb "minio/grpc/minio"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Println("Cannot listen on 6969 bitch")
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Println("Failed to serve grpc server")
	}
}
