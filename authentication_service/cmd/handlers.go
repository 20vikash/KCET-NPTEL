package main

import (
	pb "authentication/grpc/server/auth"
	"authentication/internal/gmail"
	"authentication/models"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func (a *Application) CreateUser(ctx context.Context, user *pb.UserDetails) (*pb.AuthResponse, error) {
	userD := models.User{
		Email:    user.Email,
		Password: user.Password,
		UserName: user.UserName,
	}

	ok := a.Store.Auth.CreateUser(ctx, userD)

	if !ok {
		return &pb.AuthResponse{Message: "Fail"}, errors.New("failed to create an user")
	}

	token := a.SetToken(ctx, user.Email)

	gmail.SendMail(user.Email, token)

	return &pb.AuthResponse{Message: "Created User"}, nil
}

func (a *Application) SetToken(ctx context.Context, email string) string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	t := email + time.Now().String() + fmt.Sprintf("%x", uuid)

	h := sha256.New()
	h.Write([]byte(t))
	bs := h.Sum(nil)

	token := fmt.Sprintf("%x", bs)

	a.Store.Redis.SetEmailToken(ctx, email, token)

	return token
}

func (a *Application) VerifyUser(ctx context.Context, token string) {
	value := a.Store.Redis.GetEmailFromToken(ctx, token)

	email := strings.Split(value, ":")[1]

	err := a.Store.Auth.VerifyUser(ctx, email)
	if err != nil {
		log.Println(err)
	}
}
