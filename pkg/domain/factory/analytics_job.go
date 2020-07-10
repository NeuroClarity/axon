package factory

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

// NewAnalyticsJob is a factory function to create and persist a new AnalyticsJob.
func NewAnalyticsJob(rj core.ReviewJob, db gateway.Database) (core.AnalyticsJob, error) {

	analyticsJob := core.AnalyticsJob{rid, eeg, webcam}
	err := db.NewAnalyticsJob(rid, eeg, webcam)
	if err != nil {
		return nil, err
	}

	return analyticsJob, nil
}
