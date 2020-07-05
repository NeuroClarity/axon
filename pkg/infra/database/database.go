package database

import (
	"github.com/NeuroClarity/axon/pkg/application/gateway"
	"github.com/NeuroClarity/axon/pkg/domain/core"
)

func NewDatabase(endpoint, secret string) gateway.Database {
	return database{endpoint, secret}
}

type database struct {
	databaseEndpoint string
	secret           string
}

func (repo database) GetReviewJob() core.ReviewJob {
	// database logic
	return core.ReviewJob{}
}
