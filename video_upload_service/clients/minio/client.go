package minio

import (
	"context"
	"log"
	"time"

	pb "video_upload/grpc/minio"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Upload(ctx context.Context, chunk []byte, fileName string, chunkID int64, done bool) bool {
	conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:6969: %v", err)
	}

	defer conn.Close()
	c := pb.NewFileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	chunkData := &pb.FileChunk{
		Data:     chunk,
		FileName: fileName,
		ChunkID:  chunkID,
		Done:     done,
	}

	r, err := c.UploadFile(ctx, chunkData)
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}

	return r.Success
}
