package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ReviewJobRepository interface {
	GetReviewJob() (core.ReviewJob, error)
}

func NewReviewJobRepository(database gateway.Database) ReviewJobRepository {
	return &reviewJobRepository{database}
}

type reviewJobRepository struct {
	database gateway.Database
}

func (repo *reviewJobRepository) GetReviewJob() (core.ReviewJob, error) {
	return repo.database.GetReviewJob()
}
