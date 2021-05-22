// +build debug

package main

import (
	"context"
)

func sendZoom(ctx context.Context, message string) error {
	log.Println(`[DEBUG] skip "sendZoom"`)

	return nil
}
