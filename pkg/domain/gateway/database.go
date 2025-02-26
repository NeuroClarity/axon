package gateway

import (
	"time"

	"github.com/NeuroClarity/axon/pkg/domain/core"
)

// Database gives access to a relational datastore.
type Database interface {
	NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) error
	GetReviewer(uid string) (*core.Reviewer, error)
	UpdateReviewerWithReviewJob(uid string, reviewJob *core.ReviewJob) error

	NewCreator(uid, firstName, lastName, email, company string) error
	GetCreator(uid string) (*core.Creator, error)

	NewStudy(creatorId, videoKey string, request *core.TargetAudience) (int, error)
	GetStudy(uid int) (*core.Study, error)
	GetAllStudies(creatorId string) ([]*core.Study, error)

	NewReview(reviewerId, videoKey, creatorId, eeg core.EEGData, webcam core.WebcamData) error

	GetReviewJob(demo core.Demographics, hardware core.Hardware) (*core.ReviewJob, error)
	GetReviewJobByStudy(study *core.Study) (*core.ReviewJob, error)
	NewReviewJob(studyID int, reviewerID string, completed time.Time) error
	GetStudyReviews(creatorId, videoKey string) ([]*core.Review, error)
}
