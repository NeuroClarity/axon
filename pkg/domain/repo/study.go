package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

type StudyRepository interface {
	GetStudy(uid int) (*core.Study, error)
}

func NewStudyRepository(database gateway.Database) StudyRepository {
	return &studyRepository{database}
}

type studyRepository struct {
	database gateway.Database
}

func (repo *studyRepository) GetStudy(uid int) (*core.Study, error) {
	return repo.database.GetStudy()
}
