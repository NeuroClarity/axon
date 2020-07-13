package gateway

import "time"

// Storage gives access to a filestore.
type Storage interface {
	GetRawDataUpload(dataKey string, expiration time.Duration) (string, error)
	GetVideoURL(videoKey string, expiration time.Duration) (string, error)
	GetVideoUploadURL(videoKey string, expiration time.Duration) (string, error)
	StoreBioMetricData(key, data string) error
}
