package main

import (
  "fmt"
	"testing"
  "time"
)

func TestGetSignedVideoUrl(t *testing.T) {
  storageClient := NewStorage("us-west-1")
  url, err := storageClient.getVideoUrl("testing/test-github-screenshot", time.Duration(30))
  if err != nil {
    t.Errorf(fmt.Sprintf("%s", err))
  }

  fmt.Println(url)
	return
}
