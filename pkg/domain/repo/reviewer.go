package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ReviewerRepository interface {
	GetReviewer(uid string) (*core.Reviewer, error)
	NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) error
}

func NewReviewerRepository(database gateway.Database) ReviewerRepository {
	return &reviewerRepository{database}
}

type reviewerRepository struct {
	database gateway.Database
}

func (repo *reviewerRepository) GetReviewer(uid string) (*core.Reviewer, error) {
	return &core.Reviewer{UID: "foo|123", FirstName: "John", LastName: "Smith", Email: "john@hotmail.net", Demographics: core.Demographics{Age: 50, Gender: "male", Race: "black"}}, nil
	// return repo.database.GetReviewer(uid)
}

func (repo *reviewerRepository) NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) error {
	return repo.database.NewReviewer(uid, firstName, lastName, email, demos)
}
