package gofab

import (
	"log"
	"testing"
)

func TestGOFab(t *testing.T) {
	sdk := SDK("config.yaml")
	log.Println(sdk)
}