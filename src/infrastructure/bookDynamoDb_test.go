package infrastructure

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func TestCreateItem(t *testing.T) {
	BookDynamoDb := NewBookDynamoDb(dynamodbiface.DynamoDBAPI)
}
