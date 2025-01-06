package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arpancodes/bookstore-api/handlers"
)

func RegisterRoutes(handler *handlers.BookHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", handler.ListBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handler.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id:[0-9]+}", handler.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id:[0-9]+}", handler.DeleteBook).Methods(http.MethodDelete)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Route Not Found", http.StatusNotFound)
	})

	return router
}
