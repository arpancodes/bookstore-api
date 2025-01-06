package storage

import "github.com/arpancodes/bookstore-api/models"

type InMemoryStorage struct {
	Books     []models.Book
	IDCounter int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Books:     []models.Book{},
		IDCounter: 1,
	}
}

func (s *InMemoryStorage) AddBook(book models.Book) models.Book {
	book.ID = s.IDCounter
	s.IDCounter++
	s.Books = append(s.Books, book)
	return book
}

func (s *InMemoryStorage) GetBooks() []models.Book {
	return s.Books
}

func (s *InMemoryStorage) UpdateBook(id int, updatedBook models.Book) (*models.Book, bool) {
	for i, book := range s.Books {
		if book.ID == id {
			s.Books[i] = updatedBook
			s.Books[i].ID = id
			return &s.Books[i], true
		}
	}
	return nil, false
}

func (s *InMemoryStorage) DeleteBook(id int) bool {
	for i, book := range s.Books {
		if book.ID == id {
			s.Books = append(s.Books[:i], s.Books[i+1:]...)
			return true
		}
	}
	return false
}
