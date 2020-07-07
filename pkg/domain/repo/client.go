package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type ClientRepository interface {
	GetClient() (*core.Client, error)
}

func NewClientRepository(database gateway.Database) ClientRepository {
	return &clientRepository{database}
}

type clientRepository struct {
	database gateway.Database
}

func (repo *clientRepository) GetClient() (*core.Client, error) {
	return repo.database.GetClient()
}
