package gateway

import "github.com/NeuroClarity/axon/pkg/domain/core"

type Database interface {
	GetReviewer() (*core.Reviewer, error)
	NewReviewJob() error
	GetReviewJob(age int, gender, race string) (core.ReviewJob, error)

	NewAnalyticsJob(rid string, eeg core.EEGData, webcam core.WebcamData) error
	GetAnalyticsJob(rid string, eeg core.EEGData, webcam core.WebcamData) (core.AnalyticsJob, error)

	GetStudy() (*core.Study, error)
	GetCreator() (*core.Creator, error)
	// Others...
}
