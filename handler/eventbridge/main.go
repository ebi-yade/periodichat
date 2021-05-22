package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ebi-yade/periodichat/zoom"
)

const (
	eventTypeScheduled = "Scheduled Event"
)

var (
	Revision = ""
)

func init() {
	log.Println("build revision: ", Revision)
}

func handleRequest(ctx context.Context, req events.CloudWatchEvent) error {
	this := "handleRequest"
	if req.DetailType != eventTypeScheduled {
		return fmt.Errorf(`[ERROR] in %s: invalid event type: expected="%s" but found="%s"`, this, eventTypeScheduled, req.DetailType)
	}
	eventMinute := req.Time.Unix() / 60

	items, err := getAllFromDynamoDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get job listing: %w", err)
	}

	client := zoom.New(zoomAuthToken)
	for _, item := range items {
		if int(eventMinute)%item.Divisor == 0 {
			if err := sendZoom(ctx, client, item.Message); err != nil {
				return fmt.Errorf("error sending message to zoom: %w", err)
			}
		}
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
