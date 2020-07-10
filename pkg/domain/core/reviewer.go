package core

// A Reviewer is a person who participates in studies by watching videos and providing biometric data.
type Reviewer struct {
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Demographics Demographics
}

// Demographics enumerates characteristics about a Study Participant that are interesting to a Client.
type Demographics struct {
	Age    int
	Gender string
	Race   string
}

// NewDemographics is a factory for Demographic struct. !!Need to implement socioecon!!
func NewDemographics(age int, gender, race string) Demographics {
	// TODO: Enumeration checking logic
	return Demographics{age, gender, race}
}

// UID is the primary key and unique identifier for this Reviewer.
func (r Reviewer) UID() string {
	return r.Email
}
