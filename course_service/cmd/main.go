package main

import (
	"context"
	env "course"
	pb "course/grpc/server"
	"course/internal/store"
	"log"
	"net"

	"course/internal/db"

	"google.golang.org/grpc"
)

type Application struct {
	pb.UnimplementedCourseServiceServer
	Port  string
	Store *store.Store
}

func main() {
	pg := &db.PG{
		Host:     "postgres_db",
		Username: env.GetDBUserName(),
		Password: env.GetDBPassword(),
		Database: env.GetDBName(),
	}

	db, err := pg.Connect(context.Background())
	if err != nil {
		log.Println(err)
	}

	app := &Application{
		Port:  ":5004",
		Store: store.NewStore(db),
	}

	lis, err := net.Listen("tcp", app.Port)
	if err != nil {
		log.Fatalf("failed to listen on port 5004: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, app)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
