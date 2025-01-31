package models

// Book represents a book in the bookstore
type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type Storage interface {
	AddBook(book Book) Book
	GetBooks() []Book
	UpdateBook(id int, book Book) (*Book, bool)
	DeleteBook(id int) bool
}
