package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"go-web-app/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)
	
	log.Println("Starting server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting server down")
}
