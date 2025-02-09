package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	done := r.Header.Get("done")
	isDone, _ := strconv.ParseBool(done)

	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, "Desktop", "videos", "image.mp4")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Cannot open the file")
	}
	defer file.Close()
	body := r.Body

	if isDone {
		ConvertToHls(filePath)
		log.Println("Done is true")
	}

	fmt.Println(body)
	_, err = io.Copy(file, body)
	if err != nil {
		log.Println("Error. yea.. pretty much")
	}

	w.Write([]byte("Success"))
}
