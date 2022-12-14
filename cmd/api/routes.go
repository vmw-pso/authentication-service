package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (s *server) routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s.mux.Post("/signup", s.handleSignup())
	s.mux.Post("/signin", s.handleSignin())

	return mux
}
