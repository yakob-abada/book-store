package application

import (
	"book-store/src/domain"
	"book-store/src/infrastructure"
)

type BookList struct {
	bookDynamoDb infrastructure.BookRepository
}

func NewBookList(bookDynamoDb infrastructure.BookRepository) *BookList {
	return &BookList{
		bookDynamoDb,
	}
}

func (bl *BookList) Get() (*[]domain.Book, error) {
	result, err := bl.bookDynamoDb.GetList()

	return &result, err
}
