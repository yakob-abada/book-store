package application

import (
	"book-store/src/domain"
	"book-store/src/infrastructure"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	book := domain.Book{Name: "", Featured: false, Author: ""}

	mockBookRepository := new(infrastructure.BookDynamoDbMock)

	mockBookRepository.On("Create", book).Return(nil)

	bookCreation := NewBookCreation(mockBookRepository)

	err := bookCreation.Create(book)

	assert.Nil(t, err)
}

func TestFailedCreateBook(t *testing.T) {
	book := domain.Book{Name: "", Featured: false, Author: ""}

	mockBookRepository := new(infrastructure.BookDynamoDbMock)

	mockBookRepository.On("Create", book).Return(fmt.Errorf("Something went wrong"))

	bookCreation := NewBookCreation(mockBookRepository)

	err := bookCreation.Create(book)

	assert.ErrorContains(t, err, "Something went wrong")
}
