package main

import pb "minio/grpc/minio"

type Server struct {
	pb.UnimplementedHelloWorldServiceServer
}
