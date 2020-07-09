package core

// TODO: A Study is an aggregate of users
type Study struct {
	NumParticipants int
	Demographics    Demographics
	Reviews         []Review
	ReviewJobs      []ReviewJob
	Client          Creator
	Content         Content
}

// TODO: Factory pattern... What goes into making a new study
func NewStudy(numParticipants int, demographics Demographics, client Creator, content Content) *Study {
	return &Study{NumParticipants: numParticipants, Demographics: demographics}
}
