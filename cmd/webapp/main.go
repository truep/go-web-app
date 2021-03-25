package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"

	"go-web-app/internal/config"
	"go-web-app/internal/handler"
)

func main() {
	cfg := config.Server{}
	if err := cleanenv.ReadConfig("config.yml", &cfg); err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Println("Starting server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting server down")
}
