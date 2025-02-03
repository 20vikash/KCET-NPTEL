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

	err = s.minio.uploadObject(ctx, "image1", "/tmp/image1.png")
	if err != nil {
		return &pb.HelloWorldResponse{
			Message: err.Error(),
		}, nil
	}

	return &pb.HelloWorldResponse{
		Message: "Successfully uploaded the image file",
	}, nil
}
