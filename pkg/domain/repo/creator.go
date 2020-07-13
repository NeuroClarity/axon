package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type CreatorRepository interface {
	GetCreator(uid string) (*core.Creator, error)
	NewCreator(uid, firstName, lastName, email, company string) error
}

func NewCreatorRepository(database gateway.Database) CreatorRepository {
	return &creatorRepository{database}
}

type creatorRepository struct {
	database gateway.Database
}

func (repo *creatorRepository) NewCreator(uid, firstName, lastName, email, company string) error {
  return repo.database.NewCreator(uid, firstName, lastName, email, company)
}

func (repo *creatorRepository) GetCreator(uid string) (*core.Creator, error) {
	return repo.database.GetCreator(uid)
}
