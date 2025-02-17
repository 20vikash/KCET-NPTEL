package main

import (
	pb "authentication/grpc/server/auth"
	"context"
)

type Server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func main() {

}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hello, World! "}, nil
}
