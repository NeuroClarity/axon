package repo

import (
	"time"

	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ReviewJobRepository interface {
	GetReviewJob(core.Demographics, core.Hardware) (*core.ReviewJob, error)
	GetReviewJobByStudy(study *core.Study) *core.ReviewJob
	UpdateReviewJob(reviewJob *core.ReviewJob, completed time.Time, reviewer *core.Reviewer) error
}

func NewReviewJobRepository(database gateway.Database) ReviewJobRepository {
	return &reviewJobRepository{database}
}

type reviewJobRepository struct {
	database gateway.Database
}

func (repo *reviewJobRepository) GetReviewJobByStudy(study *core.Study) *core.ReviewJob {
	return &core.ReviewJob{Study: study}
}

func (repo *reviewJobRepository) GetReviewJob(demo core.Demographics, hardware core.Hardware) (*core.ReviewJob, error) {
	return repo.database.GetReviewJob(demo, hardware)
}

// UpdateReviewJob updates a ReviewJob with completion time in persistence.
func (repo *reviewJobRepository) UpdateReviewJob(reviewJob *core.ReviewJob, completed time.Time, reviewer *core.Reviewer) error {

	studyID := reviewJob.Study.UID
	reviewerID := reviewer.UID

	return repo.database.NewReviewJob(studyID, reviewerID, completed)
}
