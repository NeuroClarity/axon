package queue

import (
  "fmt"
  "encoding/json"
  //"github.com/NeuroClarity/axon/pkg/domain/gateway"
  "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func NewQueue(sess *session.Session) (*queue, error) {
  svc := sqs.New(sess)
  return &queue{client:svc}, nil
}

type queue struct {
  client   *sqs.SQS
}

func (repo queue) PublishEegData(data string) error {
  return publishDataToQueue( , data)
}

func (repo queue) PublishEyeTrackingData(data string) error {
  return publishDataToQueue( , data)
}

func (repo queue) publishDataToQueue(url, data string) error {
  message := sqs.SendMessageInput{
    MessageBody: aws.String("message body")
  }
}
