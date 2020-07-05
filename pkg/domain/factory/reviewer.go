package factory

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

// NewReviewer is a factory function to create a new Reviewer. This should only
// occur in the event of a new Reviewer registering with NeuroClarity for the
// first time. All other successive to this entity should be made through the
// ReviewerRepository pattern.
func NewReviewer(name string, demographics core.Demographics, database gateway.Database) (*core.Reviewer, error) {
	uid, err := database.NewReviewer(name, demographics)
	if err != nil {
		return nil, err
	}
	return &core.Reviewer{uid, name, demographics}, nil
}
