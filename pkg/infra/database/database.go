package database

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

func NewDatabase(endpoint, secret string) gateway.Database {
	return database{endpoint, secret}
}

type database struct {
	databaseEndpoint string
	secret           string
}

func (repo database) NewReviewer(name string, demographics core.Demographics) (int, error) {
	// database logic
	return 0, nil
}

func (repo database) GetReviewer() (*core.Reviewer, error) {
	// database logic
	return &core.Reviewer{}, nil
}

func (repo database) GetReviewJob() (core.ReviewJob, error) {
	// database logic
	return core.ReviewJob{}, nil
}

func (repo database) NewCreator(name string) (int, error) {
	// database logic
	return 0, nil
}

func (repo database) GetCreator() (*core.Creator, error) {
	// database logic
	return &core.Creator{}, nil
}

func (repo database) GetStudy() (*core.Study, error) {
	// database logic
	return &core.Study{}, nil
}
