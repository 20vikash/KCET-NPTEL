package main

import pb "minio/grpc/minio"

type Server struct {
	pb.UnimplementedFileServiceServer
	minio MinIO
}
