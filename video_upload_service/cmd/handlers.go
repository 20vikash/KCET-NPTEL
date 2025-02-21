package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	pb "video_upload/grpc/server"
)

func (app *Application) UploadBinary(ctx context.Context, data *pb.VideoData) (*pb.Response, error) {
	done := data.Done
	isDone, _ := strconv.ParseBool(done)

	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, "Desktop", "videos", "image.mp4")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Cannot open the file")
	}
	defer file.Close()
	videoBytes := data.Data
	video := bytes.NewReader(videoBytes)

	if isDone {
		ConvertToHls(filePath)
		log.Println("Done is true")
	}

	_, err = io.Copy(file, video)
	if err != nil {
		log.Println("Error. yea.. pretty much")
	}

	return &pb.Response{Message: "Success"}, nil
}
