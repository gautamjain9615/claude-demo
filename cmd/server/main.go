// Package main is the entry point for the bookstore API server.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gautamjain9615/claude-demo/internal/handlers"
	"github.com/gautamjain9615/claude-demo/internal/middleware"
	"github.com/gautamjain9615/claude-demo/internal/store"
)

// @title			Bookstore API
// @version		1.0
// @description	A simple bookstore API built with Go and Chi.
// @host			localhost:8080
// @BasePath		/
func main() {
	bookStore := store.NewBookStore()
	bookHandler := handlers.NewBookHandler(bookStore)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Health check.
	r.Get("/health", handlers.HealthCheck)

	// Book endpoints.
	r.Route("/api/books", func(r chi.Router) {
		r.Get("/", bookHandler.ListBooks)
		r.Post("/", bookHandler.CreateBook)
		r.Get("/{id}", bookHandler.GetBook)
	})

	// Swagger UI.
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	log.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
