package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"video_upload/clients/minio"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	fileName := r.Header.Get("filename")
	chunkIdS := r.Header.Get("chunk-id")
	chunkId, err := strconv.Atoi(chunkIdS)
	if err != nil {
		log.Println("Something went wrong")
	}

	body := r.Body

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	chunk := buf.Bytes()

	minio.Upload(r.Context(), chunk, fileName, int64(chunkId))

	w.Write([]byte("Success"))
}
