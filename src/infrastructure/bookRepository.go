package infrastructure

import "book-store/src/domain"

type BookRepository interface {
	Create(item domain.Book) error
	GetList() ([]domain.Book, error)
}
