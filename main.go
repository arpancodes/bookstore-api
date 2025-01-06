package main

import (
	"fmt"
	"net/http"

	"github.com/arpancodes/bookstore-api/handlers"
	"github.com/arpancodes/bookstore-api/middlewares"
	"github.com/arpancodes/bookstore-api/routes"
	"github.com/arpancodes/bookstore-api/storage"
)

func main() {
	storage := storage.NewInMemoryStorage()
	bookHandler := handlers.NewBookHandler(storage)

	router := routes.RegisterRoutes(bookHandler)
	router.Use(middlewares.LoggingMiddleware)
	fmt.Println("Server starting on port 8000")
	http.ListenAndServe(":8000", router)
}
