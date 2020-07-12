package storage

import (
  "fmt"
  "time"
  "bytes"
  "errors"
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

func (repo storage) GetRawDataUpload(dataKey string, expiration time.Duration) (string, error) {
  return repo.getPresignedUrlForUpload(dataKey, RAW_DATA_BUCKET, expiration)
}

func (repo storage) GetVideoUrl(videoKey string, expiration time.Duration) (string, error) {
  return repo.getPresignedUrlForRetrieval(videoKey, VIDEO_BUCKET, expiration)
}

func (repo storage) GetVideoUploadURL(videoKey string, expiration time.Duration) (string, error) {
  return repo.getPresignedUrlForUpload(videoKey, VIDEO_BUCKET, expiration)
}

// key should be of the format <Type(eye-tracking/eeg)>/<username>/<video-id>
// Keeping this for now, in case we decide to pivot again
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

func (repo storage) getS3ObjectMetadata(videoKey, bucket string) (*s3.HeadObjectOutput, error) {
  output, err := repo.client.HeadObject(&s3.HeadObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(videoKey),
  })
  return output, err
}

func (repo storage) getPresignedUrlForRetrieval(key, bucket string, expiration time.Duration) (string, error) {
  _, err := repo.getS3ObjectMetadata(key, bucket)
  if err != nil {
    return "", err
  }

  req, _ := repo.client.GetObjectRequest(&s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
  })

  presignedUrl, err := req.Presign(expiration * time.Second)
  if err != nil {
    return "", err
  }

  return presignedUrl, nil
}

func (repo storage) getPresignedUrlForUpload(key, bucket string, expiration time.Duration) (string, error) {
  _, err := repo.getS3ObjectMetadata(key, bucket)
  if err != nil {
    return "", err
  }

  req, _ := repo.client.PutObjectRequest(&s3.PutObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
  })

  presignedUrl, err := req.Presign(expiration * time.Second)
  if err != nil {
    return "", err
  }

  return presignedUrl, nil
}
