package main

import (
	pb "authentication/grpc/server/auth"
	"authentication/internal/store"
	"context"
)

type Application struct {
	pb.UnimplementedHelloWorldServiceServer
	Store store.Store
}

func main() {

}

func (a *Application) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hello, World! "}, nil
}
