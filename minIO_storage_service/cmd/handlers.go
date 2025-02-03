package main

import (
	"context"
	pb "minio/grpc/minio"
)

func (s *Server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	err := s.minio.connectToMinIO()
	if err != nil {
		return &pb.HelloWorldResponse{
			Message: err.Error(),
		}, nil
	}

	return &pb.HelloWorldResponse{
		Message: "Hi brother.. I fw you heavy",
	}, nil
}
