package main

import (
	"log"
	"net/http"
)

type Application struct {
	Port string
}

func main() {
	app := &Application{
		Port: ":80",
	}

	mux := app.handleRoutes()

	err := http.ListenAndServe(app.Port, mux)
	if err != nil {
		log.Println("Something went wrong in spinning up the server")
	}
}
