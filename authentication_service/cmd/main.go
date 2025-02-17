package main

import (
	pb "authentication/grpc/server/auth"
	helper "authentication/internal"
	"authentication/internal/db"
	"authentication/internal/store"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Application struct {
	pb.UnimplementedHelloWorldServiceServer
	Port  string
	Store store.Store
}

func main() {
	pg := &db.PG{
		Host:     "localhost",
		Username: helper.GetDBUserName(),
		Password: helper.GetDBPassword(),
		Database: helper.GetDBName(),
	}

	app := &Application{
		Port:  ":5001",
		Store: pg,
	}

	lis, err := net.Listen("tcp", app.Port)
	if err != nil {
		log.Panic("Cannot start the auth gRPC server")
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &Application{})
	log.Printf("gRPC server listening at %s", app.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server auth service: %v", err)
	}
}

func (a *Application) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hello, World! "}, nil
}
