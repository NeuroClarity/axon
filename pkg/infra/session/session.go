package session

import (
  "fmt"
  "errors"
  "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession(region string) (*aws_session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
  if err != nil {
    errMsg := fmt.Sprintf("Error while creating a new AWS session, %s", err)
    return nil, errors.New(errMsg)
  }

	return &aws_session{sess: sess}, nil
}

type aws_session struct {
  sess   *session.Session
}

func (repo aws_session) GetSession() *session.Session {
  return repo.sess
}
