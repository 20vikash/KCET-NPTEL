package main

import (
	"net/http"
	"video_upload/clients/minio"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	res := minio.GetResponse()

	w.Write([]byte(res))
}
