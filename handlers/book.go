package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/arpancodes/bookstore-api/models"
	"github.com/arpancodes/bookstore-api/storage"
)

type BookHandler struct {
	Storage *storage.InMemoryStorage
}

func NewBookHandler(storage *storage.InMemoryStorage) *BookHandler {
	return &BookHandler{Storage: storage}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdBook := h.Storage.AddBook(newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books := h.Storage.GetBooks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedBookPtr, success := h.Storage.UpdateBook(id, updatedBook)
	if !success {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updatedBookPtr)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if !h.Storage.DeleteBook(id) {
		http.Error(w, "Book not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted"))
}
