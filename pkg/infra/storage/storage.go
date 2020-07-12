package storage

import (
  "fmt"
  "time"
  "bytes"
  "errors"
  //"github.com/NeuroClarity/axon/pkg/domain/gateway"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewStorage(sess *session.Session) (*storage, error) {
	svc := s3.New(sess)
	return &storage{client: svc}, nil
}

type storage struct {
  client    *s3.S3
}

const VIDEO_BUCKET = "nc-client-video-content"
const RAW_DATA_BUCKET = "nc-reviewer-raw-data"

// key should be of the format <Type(eye-tracking/eeg)>/<username>/<video-id>
// Keeping this function here, but I don't think it is going to be used w/
// current implementation
func (repo storage) StoreBioMetricData(key, data string) error {
  _, err := repo.getS3ObjectMetadata(key, RAW_DATA_BUCKET)
  if err == nil {
    return errors.New(fmt.Sprintf("Object with key %s already exists", key))
  }

  _, err = repo.client.PutObject(&s3.PutObjectInput{
    Bucket: aws.String(RAW_DATA_BUCKET),
    Key:    aws.String(key),
    ACL:    aws.String("private"),
    Body:   bytes.NewReader([]byte(data)),
  })

  if err != nil {
    return err
  }

  return nil
}

func (repo storage) GetVideoUrl(videoKey string, expiration time.Duration) (string, error) {
  _, err := repo.getS3ObjectMetadata(videoKey, VIDEO_BUCKET)
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

func (repo storage) GetVideoUploadUrl(videoKey string, expiration time.Duration) (string, error) {
  _, err := repo.getS3ObjectMetadata(videoKey, VIDEO_BUCKET)
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

func (repo storage) getS3ObjectMetadata(videoKey, bucket string) (*s3.HeadObjectOutput, error) {
  output, err := repo.client.HeadObject(&s3.HeadObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(videoKey),
  })
  return output, err
}
