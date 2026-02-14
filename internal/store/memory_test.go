package store

import (
	"testing"

	"github.com/gautamjain9615/claude-demo/internal/models"
)

func TestNewBookStore(t *testing.T) {
	s := NewBookStore()

	books := s.ListBooks()
	if len(books) != 3 {
		t.Errorf("expected 3 seeded books, got %d", len(books))
	}
}

func TestAddBook(t *testing.T) {
	s := NewBookStore()
	initialCount := len(s.ListBooks())

	book := s.AddBook(models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Price:  19.99,
	})

	if book.ID == "" {
		t.Error("expected book to have an ID assigned")
	}

	if book.Title != "Test Book" {
		t.Errorf("expected title 'Test Book', got %q", book.Title)
	}

	if len(s.ListBooks()) != initialCount+1 {
		t.Errorf("expected %d books, got %d", initialCount+1, len(s.ListBooks()))
	}
}

func TestGetBook(t *testing.T) {
	s := NewBookStore()

	tests := []struct {
		name    string
		id      string
		wantOK  bool
	}{
		{name: "existing book", id: "1", wantOK: true},
		{name: "non-existing book", id: "999", wantOK: false},
		{name: "empty id", id: "", wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := s.GetBook(tt.id)
			if ok != tt.wantOK {
				t.Errorf("GetBook(%q) ok = %v, want %v", tt.id, ok, tt.wantOK)
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	s := NewBookStore()

	tests := []struct {
		name    string
		id      string
		wantOK  bool
	}{
		{name: "existing book", id: "1", wantOK: true},
		{name: "non-existing book", id: "999", wantOK: false},
		{name: "already deleted", id: "1", wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := s.DeleteBook(tt.id)
			if ok != tt.wantOK {
				t.Errorf("DeleteBook(%q) = %v, want %v", tt.id, ok, tt.wantOK)
			}
		})
	}

	// Verify book count decreased.
	if len(s.ListBooks()) != 2 {
		t.Errorf("expected 2 books after delete, got %d", len(s.ListBooks()))
	}
}

func TestListBooks(t *testing.T) {
	s := &BookStore{
		books:  make(map[string]models.Book),
		nextID: 1,
	}

	// Empty store should return empty slice, not nil.
	books := s.ListBooks()
	if books == nil {
		t.Error("expected non-nil slice from empty store")
	}

	if len(books) != 0 {
		t.Errorf("expected 0 books, got %d", len(books))
	}
}
