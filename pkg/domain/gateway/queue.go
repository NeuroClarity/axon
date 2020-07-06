package gateway

import ()

type Queue interface {
  PublishEegData(data string) error
  PublishEyeTrackingData(data string) error
}
