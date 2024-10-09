package main

import (
	"net/http"

	"github.com/fabiobap/go-tibia-calc/internal/config"
	"github.com/fabiobap/go-tibia-calc/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/info-lvl", handlers.Repo.InfoLevel)
	mux.Post("/info-lvl", handlers.Repo.PostInfoLevel)
	mux.Get("/midnight-shards", handlers.Repo.MidnightShards)
	mux.Post("/midnight-shards", handlers.Repo.PostMidnightShards)
	mux.Get("/stone-of-insight", handlers.Repo.StoneOfInsight)
	mux.Post("/stone-of-insight", handlers.Repo.PostStoneOfInsight)

	return mux
}
