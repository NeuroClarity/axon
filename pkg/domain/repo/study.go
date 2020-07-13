package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

// URL will expire after two minutes
const URL_EXPIRATION = 120

type StudyRepository interface {
	NewStudy(creatorId, videoKey string, request *core.StudyRequest) (int, string, error)
	GetStudy(creatorId, videoKey string) (*core.Study, error)
}

func NewStudyRepository(database gateway.Database, storage gateway.Storage) *studyRepository {
	return &studyRepository{database, storage}
}

type studyRepository struct {
	database gateway.Database
	storage  gateway.Storage
}

// Returns the studyId, presignedURl to upload study content and error if applicable
func (repo *studyRepository) NewStudy(creatorId, videoKey string, request *core.StudyRequest) (int, string, error) {
	studyId, err := repo.database.NewStudy(creatorId, videoKey, request)
	if err != nil {
		return -1, "", err
	}
	url, err := repo.storage.GetVideoUploadUrl(videoKey, URL_EXPIRATION)
	return studyId, url, nil
}

func (repo *studyRepository) GetStudy(creatorId, videoKey string) (*core.Study, error) {
	return repo.database.GetStudy(creatorId, videoKey)
}
