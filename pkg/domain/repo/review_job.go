package repo

import (
	"github.com/NeuroClarity/axon/pkg/application/gateway"
	"github.com/NeuroClarity/axon/pkg/domain/core"
)

type ReviewJobRepository interface {
	GetReviewJob() core.ReviewJob
}

func NewReviewJobRepository(database gateway.Database) ReviewJobRepository {
	return reviewJobRepository{database}
}

type reviewJobRepository struct {
	database gateway.Database
}

func (repo reviewJobRepository) GetReviewJob() core.ReviewJob {
	return repo.database.GetReviewJob()
}
