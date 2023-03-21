package infrastructure

import (
	"book-store/src/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type BookDynamoDb struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewBookDynamoDb(svc *dynamodb.DynamoDB, tableName string) *BookDynamoDb {
	return &BookDynamoDb{
		svc:       svc,
		tableName: tableName,
	}
}

func (bd *BookDynamoDb) Create(item domain.Book) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(bd.tableName),
	}

	_, err = bd.svc.PutItem(input)

	return err
}

func (bd *BookDynamoDb) GetList() ([]domain.Book, error) {

	result, err := bd.svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(bd.tableName),
	})

	if err != nil {
		return nil, err
	}

	bookList := []domain.Book{}

	for _, i := range result.Items {
		item := new(domain.Book)
		err = dynamodbattribute.UnmarshalMap(i, &item)

		bookList = append(bookList, *item)

		if err != nil {
			return nil, err
		}
	}

	return bookList, nil
}
