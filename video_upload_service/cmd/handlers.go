package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	// res := minio.GetResponse()
	// fileName := r.Header.Get("file-name")

	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, "Desktop", "videos", "image.mov")

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Cannot open the file")
	}

	defer file.Close()

	body := r.Body

	_, err = io.Copy(file, body)
	if err != nil {
		log.Println("Error. yea.. pretty much")
	}

	w.Write([]byte("Success"))
}
