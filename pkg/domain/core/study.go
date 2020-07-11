package core

// TODO: A Study is an aggregate of users
type Study struct {
	UID             int
	NumParticipants int
  NumRemaining    int
	StudyRequest    StudyRequest
	Reviews         []Review
	ReviewJobs      []ReviewJob
	Creator         Creator
	Content         Content
}

type StudyRequest struct {
  MinAge      int
  MaxAge      int
  Gender      string
  Race        string
  Eeg         bool
  EyeTracking bool
}

// TODO: Factory pattern... What goes into making a new study
func NewStudy(numParticipants int, demographics Demographics, client Creator, content Content) *Study {
	return &Study{NumParticipants: numParticipants}
}
