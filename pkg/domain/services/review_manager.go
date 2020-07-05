package services

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
)

// reviewManger is internal implementation of the ReviewManager interface.
type ReviewManager struct {
	ReviewJobRepository repo.ReviewJobRepository
}

// TODO: AssignReviewTo will retrieve a ReviewJob for a specific Reviewer based
// on the pool of available ReviewJobs. The retrieved ReviewJob is sorted on:
//	* priority score relative to its Study
//	* Demographics match.between the ReviewJob and the desired Reviewer
func (rm *ReviewManager) AssignReviewTo(user *core.Reviewer) (core.ReviewJob, error) {
	return core.ReviewJob{}, nil
}

// TODO: Turns a Study into ReviewJobs
func (rm *ReviewManager) MakeReviewsFrom(study *core.Study) {
	// Generate review job for each in numParticipants
	for i := 0; i < study.NumParticipants; i++ {
		// demographic, content,
	}
}
