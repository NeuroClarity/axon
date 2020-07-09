package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type CreatorRepository interface {
	GetCreator() (*core.Creator, error)
}

func NewCreatorRepository(database gateway.Database) CreatorRepository {
	return &creatorRepository{database}
}

type creatorRepository struct {
	database gateway.Database
}

func (repo *creatorRepository) GetCreator() (*core.Creator, error) {
	return repo.database.GetCreator()
}
