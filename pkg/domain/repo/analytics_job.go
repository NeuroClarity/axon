package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type AnalyticsJobRepository interface {
	NewAnalyticsJob(*core.Biometrics) error
}

func NewAnalyticsJobRepository(database gateway.Database) AnalyticsJobRepository {
	return &analyticsJobRepository{database}
}

type analyticsJobRepository struct {
	database gateway.Database
}

func (repo *analyticsJobRepository) NewAnalyticsJob(biometrics *core.Biometrics) error {
	return nil
}
