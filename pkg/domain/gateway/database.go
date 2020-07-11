package gateway

import "github.com/NeuroClarity/axon/pkg/domain/core"

type Database interface {
	NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) (*core.Reviewer, error)
	GetReviewer(uid string) (*core.Reviewer, error)
	NewReviewJob(core.Demographics, core.Hardware) error
	GetReviewJob(core.Demographics, core.Hardware) (core.ReviewJob, error)

	NewAnalyticsJob(*core.Biometrics) error
	// GetAnalyticsJob(rid string, eeg core.EEGData, webcam core.WebcamData) (core.AnalyticsJob, error)

	GetStudy() (*core.Study, error)
	GetCreator() (*core.Creator, error)
	// Others...
}
