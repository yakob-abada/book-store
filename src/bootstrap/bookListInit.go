package bootstrap

import (
	"book-store/src/application"
	"book-store/src/infrastructure"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateBookList() *application.BookList {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return application.NewBookList(*infrastructure.NewBook(svc, "book"))
}
