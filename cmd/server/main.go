// Package main is the entry point for the bookstore API server.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
// @schemes		http https
func main() {
	bookStore := store.NewBookStore()
	bookHandler := handlers.NewBookHandler(bookStore)

	r := chi.NewRouter()

	// CORS configuration for Swagger UI on GitHub Pages.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*.github.io", "http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
