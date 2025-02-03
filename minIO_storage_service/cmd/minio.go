package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	EndPoint  string
	AccessKey string
	SecretKey string
}

func (m *MinIO) connectToMinIO() (*minio.Client, error) {
	client, err := minio.New(m.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKey, m.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Println("Cannot connect to minio server")
		return nil, err
	}

	return client, nil
}
