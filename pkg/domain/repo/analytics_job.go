package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type AnalyticsJobRepository interface {
	NewAnalyticsJob(core.ReviewJob, *core.Biometrics) error
}

func NewAnalyticsJobRepository(database gateway.Database) AnalyticsJobRepository {
	return &analyticsJobRepository{database}
}

type analyticsJobRepository struct {
	database gateway.Database
}

func (repo *analyticsJobRepository) NewAnalyticsJob(rj core.ReviewJob, biometrics *core.Biometrics) error {
	return repo.database.NewAnalyticsJob(biometrics.Reviewer.UID(), biometrics.EEGData, biometrics.WebcamData)
}
