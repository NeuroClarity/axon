package app

import (
 "github.com/NeuroClarity/axon/pkg/domain/core"
 "github.com/NeuroClarity/axon/pkg/domain/repo"
)

// creates a new study and returns the presigned URL for uploaded the content
func CreateStudy(creatorId, videoKey string, request *core.StudyRequest, studyRepo repo.StudyRepository) (int, string, error) {
  return studyRepo.NewStudy(creatorId, videoKey, request)
}

func ViewStudy(creatorId, videoKey string, studyRepo repo.StudyRepository) (*core.Study, error) {
	return studyRepo.GetStudy(creatorId, videoKey)
}

// To be implemented for the second milestone
func ListCreatorStudies(creatorId string, studyRepo repo.StudyRepository) ([]*core.Study, error) {
  return nil, nil
}
