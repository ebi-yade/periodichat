package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	Revision = ""
)

func init() {
	log.Println("build revision: ", Revision)
}

var (
	errInvalidMethod = errors.New("invalid HTTP method")
)

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod == "GET" {
		log.Println("query params:", req.QueryStringParameters)

		return events.APIGatewayProxyResponse{
			Body:       `{"success":true}`,
			StatusCode: 200,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       errInvalidMethod.Error(),
		StatusCode: http.StatusMethodNotAllowed,
	}, errInvalidMethod
}

func main() {
	lambda.Start(handleRequest)
}
