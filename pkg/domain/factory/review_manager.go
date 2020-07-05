package factory

import (
	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/domain/services"
)

// NewReviewManager is the factory function to create a new ReviewManager
// service. It should only be invoked once to enforce singleton existence of
// the ReviewManager in memory.
func NewReviewManager(jobRepo repo.ReviewJobRepository) *services.ReviewManager {
	return &services.ReviewManager{ReviewJobRepository: jobRepo}
}
