package main

import (
	"context"
	pb "minio/grpc/minio"
)

func (s *Server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "Hi brother.. I fw you heavy",
	}, nil
}
