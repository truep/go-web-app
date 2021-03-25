package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"

	"go-web-app/internal/api/jokes"
	"go-web-app/internal/config"
	"go-web-app/internal/handler"
)

func main() {
	cfg := config.Server{}
	if err := cleanenv.ReadConfig("config.yml", &cfg); err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient)
	r := chi.NewRouter()

	r.Get("/", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	log.Printf("Starting server at %s", path)
	if err := http.ListenAndServe(path, r); err != nil {
		log.Fatal(err)
	}

	log.Printf("Shutting server down")
}
