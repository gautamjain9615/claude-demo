// Package handlers provides HTTP handlers for the bookstore API.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/gautamjain9615/claude-demo/internal/models"
	"github.com/gautamjain9615/claude-demo/internal/store"
)

// BookHandler holds dependencies for book-related HTTP handlers.
type BookHandler struct {
	Store *store.BookStore
}

// NewBookHandler creates a new BookHandler.
func NewBookHandler(s *store.BookStore) *BookHandler {
	return &BookHandler{Store: s}
}

// ListBooks returns all books.
//
//	@Summary		List all books
//	@Description	Returns a list of all books in the store.
//	@Tags			books
//	@Produce		json
//	@Success		200	{array}	models.Book
//	@Router			/api/books [get]
func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books := h.Store.ListBooks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books) //nolint:errcheck // response write error is not actionable.
}

// GetBook returns a single book by ID.
//
//	@Summary		Get a book by ID
//	@Description	Returns a single book by its ID.
//	@Tags			books
//	@Produce		json
//	@Param			id	path		string	true	"Book ID"
//	@Success		200	{object}	models.Book
//	@Failure		404	{object}	ErrorResponse
//	@Router			/api/books/{id} [get]
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, ok := h.Store.GetBook(id)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "book not found"}) //nolint:errcheck // response write error is not actionable.

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book) //nolint:errcheck // response write error is not actionable.
}

// CreateBook adds a new book to the store.
//
//	@Summary		Create a new book
//	@Description	Adds a new book to the store.
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			book	body		models.Book	true	"Book to create"
//	@Success		201		{object}	models.Book
//	@Failure		400		{object}	ErrorResponse
//	@Router			/api/books [post]
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "invalid request body"}) //nolint:errcheck // response write error is not actionable.

		return
	}

	created := h.Store.AddBook(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created) //nolint:errcheck // response write error is not actionable.
}

// DeleteBook removes a book from the store by ID.
//
//	@Summary		Delete a book
//	@Description	Removes a book from the store by its ID.
//	@Tags			books
//	@Produce		json
//	@Param			id	path	string	true	"Book ID"
//	@Success		204	"No Content"
//	@Failure		404	{object}	ErrorResponse
//	@Router			/api/books/{id} [delete]
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if !h.Store.DeleteBook(id) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "book not found"}) //nolint:errcheck // response write error is not actionable.

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ErrorResponse represents an error response body.
type ErrorResponse struct {
	Message string `json:"message" example:"book not found"`
}
