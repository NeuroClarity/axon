package core

// A Study is an aggregate of users
type Study struct {
  UID             int
	Creator         *Creator
	Content         Content
	StudyRequest    *StudyRequest
	NumRemaining    int
	Reviews         []*Review
	ReviewJobs      []ReviewJob
}

// TODO: Factory pattern... What goes into making a new study
func NewStudy(creatorId, contentKey string, request *StudyRequest) *Study {
	return &Study{}
}
