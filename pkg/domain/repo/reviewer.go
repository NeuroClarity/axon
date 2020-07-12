package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ReviewerRepository interface {
	GetReviewer(uid string) (*core.Reviewer, error)
	NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) error
	AddReviewJob(uid string, job *core.ReviewJob) error
}

func NewReviewerRepository(database gateway.Database) ReviewerRepository {
	return &reviewerRepository{database}
}

type reviewerRepository struct {
	database gateway.Database
}

// GetReviewer retrieves a Reviewer from persistence.
func (repo *reviewerRepository) GetReviewer(uid string) (*core.Reviewer, error) {
	return &core.Reviewer{UID: "foo|123", FirstName: "John", LastName: "Smith", Email: "john@hotmail.net", Demographics: core.Demographics{Age: 50, Gender: "male", Race: "black"}}, nil
	// return repo.database.GetReviewer(uid)
}

// AddReviewJob adds a ReviewJob to a Reviewer to be persisted.
func (repo *reviewerRepository) AddReviewJob(uid string, job *core.ReviewJob) error {
	// TODO: talk about db implementation
	return repo.database.UpdateReviewerWithReviewJob(uid, job)
}

// NewReviewer creates a Reviewer in persistence.
func (repo *reviewerRepository) NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) error {
	return repo.database.NewReviewer(uid, firstName, lastName, email, demos)
}
