// Package store provides in-memory data storage for the bookstore API.
package store

import (
	"fmt"
	"sync"

	"github.com/gautamjain9615/claude-demo/internal/models"
)

// BookStore provides in-memory storage for books.
type BookStore struct {
	mu     sync.RWMutex
	books  map[string]models.Book
	nextID int
}

// NewBookStore creates a new in-memory book store with sample data.
func NewBookStore() *BookStore {
	s := &BookStore{
		books:  make(map[string]models.Book),
		nextID: 1,
	}

	// Seed with sample data.
	s.AddBook(models.Book{Title: "The Go Programming Language", Author: "Alan Donovan", Price: 39.99})
	s.AddBook(models.Book{Title: "Designing Data-Intensive Applications", Author: "Martin Kleppmann", Price: 44.99})
	s.AddBook(models.Book{Title: "Clean Code", Author: "Robert C. Martin", Price: 34.99})

	return s
}

// ListBooks returns all books in the store.
func (s *BookStore) ListBooks() []models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	books := make([]models.Book, 0, len(s.books))
	for _, b := range s.books {
		books = append(books, b)
	}

	return books
}

// GetBook returns a book by ID.
func (s *BookStore) GetBook(id string) (models.Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, ok := s.books[id]

	return book, ok
}

// AddBook adds a new book to the store and returns it with an assigned ID.
func (s *BookStore) AddBook(book models.Book) models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	book.ID = fmt.Sprintf("%d", s.nextID)
	s.nextID++
	s.books[book.ID] = book

	return book
}

// DeleteBook removes a book by ID and returns whether it existed.
func (s *BookStore) DeleteBook(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.books[id]
	if !ok {
		return false
	}

	delete(s.books, id)

	return true
}
