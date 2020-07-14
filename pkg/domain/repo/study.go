package repo

import (
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/gateway"
)

// URL will expire after two minutes
const URL_EXPIRATION = 120

type StudyRepository interface {
	NewStudy(creatorId, videoKey string, request *core.StudyRequest) (int, string, error)
	GetStudy(studyID int) (*core.Study, error)
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
	url, err := repo.storage.GetVideoUploadURL(videoKey, URL_EXPIRATION)
	if err != nil {
		return -1, "", err
	}
	studyId, err := repo.database.NewStudy(creatorId, videoKey, request)
	return studyId, url, nil
}

func (repo *studyRepository) GetStudy(studyID int) (*core.Study, error) {
	return repo.database.GetStudy(studyID)
}
