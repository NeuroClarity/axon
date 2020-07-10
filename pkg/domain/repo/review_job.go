package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ReviewJobRepository interface {
	GetReviewJob(core.Demographics) (core.ReviewJob, error)
}

func NewReviewJobRepository(database gateway.Database) ReviewJobRepository {
	return &reviewJobRepository{database}
}

type reviewJobRepository struct {
	database gateway.Database
}

func (repo *reviewJobRepository) GetReviewJob(demo core.Demographics) (core.ReviewJob, error) {
	age := demo.Age
	gender := demo.Gender
	race := demo.Race
	return repo.database.GetReviewJob(age, gender, race)
}
