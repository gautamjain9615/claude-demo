package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/gautamjain9615/claude-demo/internal/models"
	"github.com/gautamjain9615/claude-demo/internal/store"
)

func setupRouter() (*chi.Mux, *BookHandler) {
	s := store.NewBookStore()
	h := NewBookHandler(s)

	r := chi.NewRouter()
	r.Get("/api/books", h.ListBooks)
	r.Post("/api/books", h.CreateBook)
	r.Get("/api/books/{id}", h.GetBook)

	return r, h
}

func TestListBooks(t *testing.T) {
	r, _ := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/api/books", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var books []models.Book
	if err := json.Unmarshal(w.Body.Bytes(), &books); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(books) != 3 {
		t.Errorf("expected 3 books, got %d", len(books))
	}
}

func TestGetBook(t *testing.T) {
	r, _ := setupRouter()

	tests := []struct {
		name       string
		id         string
		wantStatus int
	}{
		{name: "existing book", id: "1", wantStatus: http.StatusOK},
		{name: "non-existing book", id: "999", wantStatus: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/books/"+tt.id, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	r, _ := setupRouter()

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{
			name:       "valid book",
			body:       `{"title":"New Book","author":"New Author","price":29.99}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "invalid json",
			body:       `{invalid}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/books", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusCreated {
				var book models.Book
				if err := json.Unmarshal(w.Body.Bytes(), &book); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if book.ID == "" {
					t.Error("expected book to have an ID")
				}
			}
		})
	}
}

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", http.NoBody)
	w := httptest.NewRecorder()

	HealthCheck(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var resp HealthResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status 'ok', got %q", resp.Status)
	}
}
