package storage

import (
  "fmt"
	"testing"
  "time"

)


func TestGetSignedVideoUrl(t *testing.T) {
  storageClient, _ := NewStorage("us-west-1")
  url, err := storageClient.GetVideoUrl("testing/test-github-screenshot", time.Duration(30))
  if err != nil {
    t.Errorf(fmt.Sprintf("%s", err))
  }

  fmt.Println(url)
	return
}

func TestUploadBioMetricData(t *testing.T) {
  storageClient, _ := NewStorage("us-west-1")
  data := "hello world! is this working?"
  err := storageClient.StoreBioMetricData("testing/random-text", data)
  if err != nil {
    t.Errorf(fmt.Sprintf("Error uploading data: %s", err))
  }
  return
}
