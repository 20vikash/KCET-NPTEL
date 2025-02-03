package main

import (
	"fmt"
	"net/http"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	// res := minio.GetResponse()
	// fileName := r.Header.Get("file-name")
	body := r.Body

	fmt.Println(body)

	// w.Write([]byte(res))
}
