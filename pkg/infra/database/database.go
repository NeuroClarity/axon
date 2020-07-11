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

func (repo database) GetReviewer(uid string) (*core.Reviewer, error) {
	// database logic
	return &core.Reviewer{}, nil
}

func (repo database) NewReviewer(uid, firstName, lastName, email string, demos core.Demographics) (*core.Reviewer, error) {
	return &core.Reviewer{}, nil
}

func (repo database) NewReviewJob(demo core.Demographics, hardware core.Hardware) error {
	return nil
}

func (repo database) GetReviewJob(demo core.Demographics, hardware core.Hardware) (core.ReviewJob, error) {
	// database logic
	return core.ReviewJob{}, nil
}

func (repo database) GetCreator() (*core.Creator, error) {
	// database logic
	return &core.Creator{}, nil
}

func (repo database) GetStudy() (*core.Study, error) {
	// database logic
	return &core.Study{}, nil
}

func (repo database) GetAnalyticsJob(rid string, egg core.EEGData, webcam core.WebcamData) (core.AnalyticsJob, error) {
	// database logic
	return core.AnalyticsJob{}, nil
}

func (repo database) NewAnalyticsJob(*core.Biometrics) error {
	return nil
}
