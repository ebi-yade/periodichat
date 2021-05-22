// +build !debug

package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/ebi-yade/periodichat/env"
)

type JobItem struct {
	UUID    string
	Divisor int
	Message string
}

var (
	tableName           = env.MustNonEmpty("TABLE_NAME")
	dynamodbClient      = mustNewClient()
	zeroValueScanOutput = make([]JobItem, 0)

	errInvalidDynamoDBItem = errors.New("invalid DynamoDB item")
)

func mustNewClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	return dynamodb.NewFromConfig(cfg)
}

func getAllFromDynamoDB(ctx context.Context) ([]JobItem, error) {
	out, err := dynamodbClient.Scan(ctx, &(dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}))

	if err != nil {
		return zeroValueScanOutput, err
	}

	items := make([]JobItem, 0, cap(out.Items))
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &items); err != nil {
		return zeroValueScanOutput, fmt.Errorf(`%w: %s`, errInvalidDynamoDBItem, err.Error())
	}

	for _, item := range items {
		if item.Divisor < 0 {
			return zeroValueScanOutput, fmt.Errorf(`%w: Divisor must be more than 0 but was %d`, errInvalidDynamoDBItem, item.Divisor)
		}
	}

	return items, nil
}
