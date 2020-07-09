package queue

import (
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

const eegDataQueueUrl = "https://sqs.us-west-1.amazonaws.com/471943556279/nc-eeg-data"
const eyeDataQueueUrl = "https://sqs.us-west-1.amazonaws.com/471943556279/nc-eye-tracking-data"

func (repo queue) PublishEegDataKey(key string) error {
  return repo.publishToQueue(eegDataQueueUrl, key)
}

func (repo queue) PublishEyeTrackingDataKey(key string) error {
  return repo.publishToQueue(eyeDataQueueUrl, key)
}

func (repo queue) publishToQueue(url, data string) error {
  _, err := repo.client.SendMessage(&sqs.SendMessageInput{
    MessageAttributes: map[string]*sqs.MessageAttributeValue{
      "Key": &sqs.MessageAttributeValue{
        DataType:     aws.String("String"),
        StringValue:  aws.String(data),
      },
    },
    MessageBody: aws.String("S3 Key"),
    QueueUrl: &url,
  })

  return err
}
