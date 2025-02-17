package main

import (
	pb "authentication/grpc/server/auth"
	helper "authentication/internal"
	"authentication/internal/db"
	"authentication/internal/store"
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Application struct {
	pb.UnimplementedAuthServiceServer
	Port  string
	Store *store.Store
}

func main() {
	pg := &db.PG{
		Host:     "localhost",
		Username: helper.GetDBUserName(),
		Password: helper.GetDBPassword(),
		Database: helper.GetDBName(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := pg.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}

	app := &Application{
		Port:  ":5001",
		Store: store.NewStore(conn),
	}

	lis, err := net.Listen("tcp", app.Port)
	if err != nil {
		log.Panic("Cannot start the auth gRPC server")
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, app)
	log.Printf("gRPC server listening at %s", app.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server auth service: %v", err)
	}
}
