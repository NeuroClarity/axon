package core

// A Reviewer is a person who participates in studies by watching videos and providing biometric data.
type Reviewer struct {
	UID          string
	FirstName    string
	LastName     string
	Email        string
	Demographics Demographics
	ReviewJobs   []ReviewJob
}

// Demographics enumerates characteristics about a Study Participant that are interesting to a Client.
type Demographics struct {
	Age    int
	Gender string
	Race   string
}
