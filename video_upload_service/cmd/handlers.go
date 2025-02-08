package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, "Desktop", "videos", "image.png")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Cannot open the file")
	}
	defer file.Close()
	body := r.Body

	fmt.Println(body)
	_, err = io.Copy(file, body)
	if err != nil {
		log.Println("Error. yea.. pretty much")
	}

	w.Write([]byte("Success"))
}
