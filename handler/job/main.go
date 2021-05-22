package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ebi-yade/periodichat/zoom"
)

var (
	Revision = ""
)

func init() {
	log.Println("build revision: ", Revision)
}

var (
	errMalformedRequestBody = errors.New("malformed request body")
)

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) error {
	this := "handleRequest"
	log.Println("method:", req.HTTPMethod)
	log.Println("method:", req.Headers)

	var body zoom.BotNotification
	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		return fmt.Errorf(`[ERROR] in %s: %w: failed to parse json: %s`, this, errMalformedRequestBody, err.Error())
	}
	rawCmd := body.Payload.Cmd
	commands := strings.Fields(rawCmd)
	if len(commands) == 0 {
		return fmt.Errorf(`[ERROR] in %s: %w: command is empty`, this, errMalformedRequestBody)
	}

	switch commands[0] {
	case "register":
		if err := register(ctx, rawCmd); err != nil {
			return fmt.Errorf(`[ERROR] in register: %w`, err)
		}
	case "cancel":
		if err := cancel(ctx); err != nil {
			return fmt.Errorf(`[ERROR] in cancel: %w`, err)
		}
	default:
		return fmt.Errorf(`[ERROR] in %s: the command "%s" is not defined: %w`, this, commands[0], errMalformedRequestBody)
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
