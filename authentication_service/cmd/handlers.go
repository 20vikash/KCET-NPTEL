package main

import (
	pb "authentication/grpc/server/auth"
	"authentication/models"
	"context"
	"errors"
)

func (a *Application) CreateUser(ctx context.Context, user *pb.UserDetails) (*pb.AuthResponse, error) {
	userD := models.User{
		Email:    user.Email,
		Password: user.Password,
		UserName: user.UserName,
	}

	ok := a.Store.Auth.CreateUser(ctx, userD)

	if ok {
		return &pb.AuthResponse{Message: "Created User"}, nil
	}

	return &pb.AuthResponse{Message: "Fail"}, errors.New("failed to create an user")
}
