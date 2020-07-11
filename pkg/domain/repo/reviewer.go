package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ReviewerRepository interface {
	GetReviewer(uid string) (*core.Reviewer, error)
	NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) (*core.Reviewer, error)
}

func NewReviewerRepository(database gateway.Database) ReviewerRepository {
	return &reviewerRepository{database}
}

type reviewerRepository struct {
	database gateway.Database
}

func (repo *reviewerRepository) GetReviewer(uid string) (*core.Reviewer, error) {
	return nil, nil
	// return repo.database.GetReviewer(uid)
}

func (repo *reviewerRepository) NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) (*core.Reviewer, error) {
	return repo.database.NewReviewer(uid, firstName, lastName, email, demos)
}
