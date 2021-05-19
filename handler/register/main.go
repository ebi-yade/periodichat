package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	Revision = ""
)

func init() {
	log.Println("build revision: ", Revision)
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("method:", req.HTTPMethod)
	log.Println("path:", req.Path)
	log.Println("body:", req.Body)

	log.Println("path params:", req.PathParameters)
	log.Println("query params:", req.QueryStringParameters)

	return events.APIGatewayProxyResponse{
		Body:       "request received!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
