package infrastructure

import (
	"book-store/src/domain"

	"github.com/stretchr/testify/mock"
)

type BookDynamoDbMock struct {
	mock.Mock
}

func (bdm *BookDynamoDbMock) Create(item domain.Book) error {
	args := bdm.Called(item)

	return args.Error(0)
}

func (bdm *BookDynamoDbMock) GetList() ([]domain.Book, error) {
	args := bdm.Called()

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain.Book), args.Error(1)
}
