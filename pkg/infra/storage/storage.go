package main

import (
  "fmt"
  "time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewStorage(region string) *storage {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
  if err != nil {
    fmt.Printf("Error while creating a new AWS session, %s", err)
  }

	svc := s3.New(sess)
	return &storage{svc}
}

type storage struct {
  client    *s3.S3
}

const VIDEO_BUCKET = "client-video-content"

func (repo storage) getVideoUrl(videoKey string, expiration time.Duration) (string, error) {
  _, err := repo.client.HeadObject(&s3.HeadObjectInput{
    Bucket: aws.String(VIDEO_BUCKET),
    Key:    aws.String(videoKey),
  })
  if err != nil {
    return "", err
  }

  req, _ := repo.client.GetObjectRequest(&s3.GetObjectInput{
    Bucket: aws.String(VIDEO_BUCKET),
    Key:    aws.String(videoKey),
  })

  presignedUrl, err := req.Presign(expiration * time.Second)
  if err != nil {
    return "", err
  }

  return presignedUrl, nil
}

func (repo storage) getVideoUploadURL() string {
  _, err := repo.client.HeadObject(&s3.HeadObjectInput{
    Bucket: aws.String(VIDEO_BUCKET),
    Key:    aws.String(videoKey),
  })
  if err != nil {
    return "", err
  }

  req, _ := repo.client.GetObjectRequest(&s3.GetObjectInput{
    Bucket: aws.String(VIDEO_BUCKET),
    Key:    aws.String(videoKey),
  })

  presignedUrl, err := req.Presign(expiration * time.Second)
  if err != nil {
    return "", err
  }

  return presignedUrl, nil
}
