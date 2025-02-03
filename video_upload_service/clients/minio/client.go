package minio

import (
	"context"
	"log"
	"time"

	pb "video_upload/grpc/minio"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetResponse() string {
	conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}

	defer conn.Close()
	c := pb.NewHelloWorldServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloWorldRequest{})
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}

	return r.GetMessage()
}
