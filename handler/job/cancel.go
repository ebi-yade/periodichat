package main

import (
	"context"
	"fmt"
)

func cancel(ctx context.Context) error {
	out, err := getAllFromDynamoDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get job listing: %w", err)
	}

	if err = deleteAllOfDynamoDB(ctx, out.Items); err != nil {
		return fmt.Errorf("failed to send delete requests: %w", err)
	}

	return nil
}
