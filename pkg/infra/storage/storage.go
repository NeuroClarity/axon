package storage

import (
  "fmt"
  "time"
  "github.com/NeuroClarity/axon/pkg/domain/gateway"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewStorage(region string) (*gateway.Storage, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
  if err != nil {
    errMsg := fmt.Printf("Error while creating a new AWS session, %s", err)
    return nil, errors.New(errMsg);
  }

	svc := s3.New(sess)
	return &storage{
          client: svc
        }, nil
}

type storage struct {
  client    *s3.S3
}

const VIDEO_BUCKET = "client-video-content"

func (repo storage) GetVideoUrl(videoKey string, expiration time.Duration) (string, error) {
  _, err := getS3ObjectMetadata(videoKey)
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

func (repo storage) GetVideoUploadURL(videoKey string, expiration time.Duration) (string, error) {
  _, err := getS3ObjectMetadata(videoKey)
  if err == nil {
    return "", errors.New(fmt.Sprintf("Object with key %s already exists", videoKey))
  }

  req, _ := repo.client.PutObjectRequest(&s3.PutObjectInput{
    Bucket: aws.String(VIDEO_BUCKET),
    Key:    aws.String(videoKey),
  })

  presignedUrl, err := req.Presign(expiration * time.Second)
  if err != nil {
    return "", err
  }

  return presignedUrl, nil
}

func getS3ObjectMetadata(videoKey string) (*s3.HeadObjectOutput, error) {
  output, err := repo.client.HeadObject(&s3.HeadObjectInput{
    Bucket: aws.String(VIDEO_BUCKET),
    Key:    aws.String(videoKey),
  })
  return output, err
}
