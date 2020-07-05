package gateway

import "github.com/NeuroClarity/axon/pkg/domain/core"

type Database interface {
	GetReviewJob() core.ReviewJob
	NewReviewer(name string, demographics core.Demographics) (int, error)
	GetReviewer() core.Reviewer
	// Others...
}
