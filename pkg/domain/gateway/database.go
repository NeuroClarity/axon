package gateway

import "github.com/NeuroClarity/axon/pkg/domain/core"

type Database interface {
	NewReviewer(name string, demographics core.Demographics) (int, error)
	GetReviewer() (*core.Reviewer, error)
	GetReviewJob() (core.ReviewJob, error)
	NewClient(name string) (int, error)
	GetClient() (*core.Client, error)
	GetStudy() (*core.Study, error)
	// Others...
}
