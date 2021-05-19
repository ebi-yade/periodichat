package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

var (
	errNoActionSpecified = errors.New("action is not specified as a part of request path")
	errActionNotFound    = errors.New("action is not found")

	Revision = ""
)

func init() {
	fmt.Println("build revision: ", Revision)
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	method := req.HTTPMethod
	body := req.Body
	action, ok := req.PathParameters["action"]
	if !ok {
		return events.APIGatewayProxyResponse{
			Body:       errNoActionSpecified.Error(),
			StatusCode: 500,
		}, errNoActionSpecified
	}

	switch action {
	case "register":
		log.Println(method)
		log.Println(body)

		return events.APIGatewayProxyResponse{
			Body:       "request received!",
			StatusCode: 200,
		}, nil
	default:

		return events.APIGatewayProxyResponse{
			Body:       errActionNotFound.Error(),
			StatusCode: 500,
		}, errActionNotFound
	}
}

func main() {
	lambda.Start(handleRequest)
}
