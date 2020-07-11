package gateway

import "github.com/NeuroClarity/axon/pkg/domain/core"

type Database interface {
	NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) error
	GetReviewer(uid string) (*core.Reviewer, error)

	NewCreator(uid, firstName, lastName, email, company string) error
	GetCreator(uid string) (*core.Creator, error)

	NewStudy(creatorId int, videoKey string, reviewCount, agemax, ageMin int, gender, race string, eegHeadset, eyeTracking bool) error
	GetStudy(creatorId, videoKey string) (*core.Study, error)
	GetAllStudies(creatorId string) ([]*core.Study, error)

	NewReview(reviewerId, videoKey, creatorId, eeg core.EEGData, webcam core.WebcamData) error

	// NewReviewJob(core.Demographics, core.Hardware) error
	// GetReviewJob(core.Demographics, core.Hardware) (core.ReviewJob, error)

	// NewAnalyticsJob(*core.Biometrics) error
	// GetAnalyticsJob(rid string, eeg core.EEGData, webcam core.WebcamData) (core.AnalyticsJob, error)
}
