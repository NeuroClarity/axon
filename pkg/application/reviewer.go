package app

import (
	"fmt"

	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
)

// InitReviewer persists Reviewer information the first time he/she logs in.
func InitReviewer() {

}

// AssignReviewJob retrieves a ReviewJob for a Reviewer.
func AssignReviewJob(reviewer *core.Reviewer, hardware core.Hardware, rjRepo repo.ReviewJobRepository) (*core.ReviewJob, error) {
	fmt.Printf("\nreviewer: %+v\n", reviewer)
	fmt.Printf("hardware: %+v\n", hardware)
	return rjRepo.GetReviewJob(reviewer.Demographics, hardware)
}

// SubmitAnalyticsJob turns a ReviewJob into an AnalyticsJob to process into Insights.
func SubmitAnalyticsJob(biometrics *core.Biometrics, ajRepo repo.AnalyticsJobRepository) error {
	return ajRepo.NewAnalyticsJob(biometrics)
}
