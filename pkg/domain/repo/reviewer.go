package repo

import (
	"github.com/NeuroClarity/axon/pkg/application/gateway"
	"github.com/NeuroClarity/axon/pkg/domain/core"
)

type ReviewerRepository interface {
	GetReviewer(uid int) *core.Reviewer
}

func NewReviewerRepository(database gateway.Database) ReviewerRepository {
	return reviewerRepository{database}
}

type reviewerRepository struct {
	database gateway.Database
}

func (repo reviewerRepository) GetReviewer(uid int) *core.Reviewer {
	return repo.database.GetReviewer()
}
