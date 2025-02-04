package main

import (
	"bytes"
	"context"
	"fmt"
	pb "minio/grpc/minio"

	"github.com/minio/minio-go/v7"
)

func (s *Server) UploadFile(ctx context.Context, chunk *pb.FileChunk) (*pb.UploadStatus, error) {
	err := s.minio.connectToMinIO()
	if err != nil {
		return &pb.UploadStatus{
			Success: false,
		}, err
	}

	reader := bytes.NewReader(chunk.Data)

	_, err = s.minio.Connection.PutObject(
		ctx,
		"videos",
		fmt.Sprintf("%s%d", chunk.FileName, chunk.ChunkID),
		reader,
		int64(len(chunk.Data)),
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)
	if err != nil {
		return &pb.UploadStatus{
			Success: false,
		}, err
	}

	return &pb.UploadStatus{
		Success: true,
	}, nil
}
