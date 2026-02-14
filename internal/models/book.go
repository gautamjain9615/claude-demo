// Package models defines data structures used across the bookstore API.
package models

import "time"

// Book represents a book in the bookstore.
type Book struct {
	ID     string `json:"id" example:"1"`
	Title  string `json:"title" example:"The Go Programming Language"`
	Author string `json:"author" example:"Alan Donovan"`
	Price  float64 `json:"price" example:"39.99"`
}

// Review represents a review for a book.
type Review struct {
	ID           string    `json:"id" example:"1"`
	BookID       string    `json:"book_id" example:"1"`
	ReviewerName string    `json:"reviewer_name" example:"Jane Doe"`
	Rating       int       `json:"rating" example:"5"`
	Comment      string    `json:"comment" example:"Excellent book!"`
	CreatedAt    time.Time `json:"created_at" example:"2024-01-15T10:30:00Z"`
}
