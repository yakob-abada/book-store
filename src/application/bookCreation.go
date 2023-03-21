package application

import (
	"book-store/src/domain"
	"book-store/src/infrastructure"
)

type BookCreation struct {
	bookDynamoDb infrastructure.BookRepository
}

func NewBookCreation(bookDynamoDb infrastructure.BookRepository) *BookCreation {
	return &BookCreation{
		bookDynamoDb,
	}
}

func (bc *BookCreation) Create(item domain.Book) error {
	return bc.bookDynamoDb.Create(item)
}
