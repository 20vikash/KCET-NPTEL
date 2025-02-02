package main

import "net/http"

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello stupid ass"))
}
