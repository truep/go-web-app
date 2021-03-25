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

	path := cfg.Host + ":" + cfg.Port

	log.Printf("Starting server at %s", path)
	if err := http.ListenAndServe(path, r); err != nil {
		log.Fatal(err)
	}

	log.Printf("Shutting server down")
}
