package gateway

import (
  "time"
)

type Storage interface {
  GetVideoUrl(videoKey string, expiration time.Duration) (string, error)
  GetVideoUploadUrl(videoKey string, expiration time.Duration) (string, error)
}
