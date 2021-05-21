// +build !debug

package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"

	"github.com/ebi-yade/periodichat/env"
)

var (
	tableName      = env.MustNonEmpty("TABLE_NAME")
	dynamodbClient = mustNewClient()

	errUUID = errors.New("UUID error")
)

func mustNewClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	return dynamodb.NewFromConfig(cfg)
}

func registerToDynamoDB(ctx context.Context, div int, msg string) error {
	randomID, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to create UUID: %w", errUUID)
	}

	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"UUID":    &types.AttributeValueMemberS{Value: randomID.String()},
			"Divisor": &types.AttributeValueMemberN{Value: strconv.Itoa(div)},
			"Message": &types.AttributeValueMemberS{Value: msg},
		},
		TableName:                aws.String(tableName),
		ConditionExpression:      aws.String("attribute_not_exists(#uuid)"),
		ExpressionAttributeNames: map[string]string{"#uuid": "UUID"},
	}

	if _, err := dynamodbClient.PutItem(ctx, input); err != nil {
		var cce *types.ConditionalCheckFailedException
		if errors.As(err, &cce) {
			return fmt.Errorf("UUID collision: %w\nfailed to register a message to DynamoDB: %s", errUUID, err.Error())
		}

		return fmt.Errorf("failed to register a message to DynamoDB: %w", err)
	}

	return nil
}
