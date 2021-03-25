package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	h := handler.NewHandler(apiClient, cfg.CustomJoke)
	r := chi.NewRouter()

	r.Get("/", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	srv := &http.Server{
		Addr:    path,
		Handler: r,
	}

	// handle shutdown gracefully
	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		// here all cleanup, close db conn and so on...
		err := srv.Shutdown(ctx)
		// ... other func
		done <- err
	}()

	log.Printf("Starting server at %s", path)
	_ = srv.ListenAndServe()

	err := <-done

	log.Printf("Shutting server down with %v", err)
}
