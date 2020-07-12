package app

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
)

// InitReviewer persists Reviewer information the first time he/she logs in.
func InitReviewer() {

}

// AssignReviewJob retrieves a ReviewJob for a Reviewer.
func AssignReviewJob(reviewer *core.Reviewer, hardware core.Hardware, rjRepo repo.ReviewJobRepository) (*core.ReviewJob, error) {
	return rjRepo.GetReviewJob(reviewer.Demographics, hardware)
}

// FinishReviewJob completes a ReviewJob and creates an AnalyticsJob.
func FinishReviewJob(reviewer *core.Reviewer, study *core.Study, biometrics *core.Biometrics, rjRepo repo.ReviewJobRepository, rRepo repo.ReviewerRepository, ajRepo repo.AnalyticsJobRepository) error {

	// Two things we have to worry about:

	//		* Assigning a completed ReviewJob to Reviewer
	reviewJob := rjRepo.GetReviewJobByStudy(study)
	reviewJob.Completed = biometrics.Time
	if err := rRepo.AddReviewJob(reviewer.UID, reviewJob); err != nil {
		return err
	}

	//		* Creating a new AnalyticsJob
	if err := ajRepo.NewAnalyticsJob(biometrics); err != nil {
		return err
	}

	return nil
}

// SubmitAnalyticsJob turns a ReviewJob into an AnalyticsJob to process into Insights.
func SubmitAnalyticsJob(biometrics *core.Biometrics, ajRepo repo.AnalyticsJobRepository) error {
	return ajRepo.NewAnalyticsJob(biometrics)
}
