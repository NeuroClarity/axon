package app

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
)

// AssignReviewJob retrieves a ReviewJob for a Reviewer.
func AssignReviewJob(reviewer *core.Reviewer, rjRepo repo.ReviewJobRepository) (core.ReviewJob, error) {
	return rjRepo.GetReviewJob(reviewer.Demographics)
}

// SubmitAnalyticsJob turns a ReviewJob into an AnalyticsJob to process into Insights.
func SubmitAnalyticsJob(reviewJob core.ReviewJob, biometrics *core.Biometrics, ajRepo repo.AnalyticsJobRepository) error {
	return ajRepo.NewAnalyticsJob(reviewJob, biometrics)
}
