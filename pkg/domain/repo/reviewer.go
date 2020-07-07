package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ReviewerRepository interface {
	GetReviewer(uid int) (*core.Reviewer, error)
}

func NewReviewerRepository(database gateway.Database) ReviewerRepository {
	return &reviewerRepository{database}
}

type reviewerRepository struct {
	database gateway.Database
}

func (repo *reviewerRepository) GetReviewer(uid int) (*core.Reviewer, error) {
	return repo.database.GetReviewer()
}
