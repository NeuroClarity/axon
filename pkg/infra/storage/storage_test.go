package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/NeuroClarity/axon/pkg/infra/session"
)

func TestGetSignedVideoUrl(t *testing.T) {
	sess, _ := session.NewSession("us-west-1")
	storageClient, err := NewStorage(sess.GetSession())
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
		return
	}
	url, err := storageClient.GetVideoUploadURL("dev-testing/sample-screenshot", time.Duration(80))
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
		return
	}

	fmt.Println(url)
	return
}

//func TestUploadBioMetricData(t *testing.T) {
//sess, _ := session.NewSession("us-west-1")
//storageClient, err := NewStorage(sess.GetSession())
//if err != nil {
//t.Errorf(fmt.Sprintf("%s", err))
//return
//}
//data := "hello world! is this working?"
//err = storageClient.StoreBioMetricData("testing/random-text", data)
//if err != nil {
//t.Errorf(fmt.Sprintf("Error uploading data: %s", err))
//return
//}
//return
//}
