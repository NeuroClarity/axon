package queue

import (
  "testing"
  "fmt"
  "github.com/NeuroClarity/axon/pkg/infra/session"
)

func TestPublishEegDataKey(t *testing.T) {
  sess, _ := session.NewSession("us-west-1")
  client, err := NewQueue(sess.GetSession())
  if err != nil {
    t.Errorf(fmt.Sprintf("%s", err))
    return
  }

  err = client.PublishEegDataKey("testing/testkey")
  if err != nil {
    t.Errorf(fmt.Sprintf("%s", err))
    return
  }

  fmt.Println("Successfully published data to queue")
  return
}
