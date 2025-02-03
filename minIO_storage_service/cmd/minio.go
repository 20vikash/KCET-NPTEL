package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	Connection *minio.Client
	EndPoint   string
	AccessKey  string
	SecretKey  string
}

func (m *MinIO) connectToMinIO() error {
	client, err := minio.New(m.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKey, m.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Println("Cannot connect to minio server")
		return err
	}

	m.Connection = client

	return nil
}

func (m *MinIO) uploadObject(ctx context.Context, objectName, filePath string) error {
	info, err := m.Connection.FPutObject(ctx, "videos", objectName, filePath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return nil
}
