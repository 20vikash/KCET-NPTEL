package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Port    string
	Handler *chi.Mux
}

func main() {
	app := &Application{
		Port:    ":8080",
		Handler: handleRoutes(),
	}

	err := http.ListenAndServe(app.Port, app.Handler)
	if err != nil {
		log.Println("Something went wrong in spinning up the server")
	}
}
