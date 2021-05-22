// +build !debug

package main

import (
	"context"
	"log"

	"github.com/ebi-yade/periodichat/env"
	"github.com/ebi-yade/periodichat/zoom"
)

var (
	zoomAuthToken = env.MustNonEmpty("ZOOM_AUTH_TOKEN")
)

func sendZoom(ctx context.Context, z *zoom.Client, msg string) error {
	log.Println("sending message...", msg)
	return nil
}
