package main

import (
	"book-store/src/bootstrap"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	BookList, err := bootstrap.CreateBookList().Get()

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	b, err := json.Marshal(BookList)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%v", string(b)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
