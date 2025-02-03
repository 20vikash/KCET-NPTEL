package main

import "os"

func loadMinIOAccessKey() string {
	return os.Getenv("MINIO_ACCESS_KEY")
}

func loadMinIOSecretKey() string {
	return os.Getenv("MINIO_SECRET_KEY")
}
