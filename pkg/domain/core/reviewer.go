package core

// A Reviewer is a person who participates in studies by watching videos and providing biometric data.
type Reviewer struct {
	UID          int
	Username     string
	FirstName    string
	LastName     string
	Password     string
	Token        string
	Email        string
	Demographics Demographics
}

// Demographics enumerates characteristics about a Study Participant that are interesting to a Client.
type Demographics struct {
	Age       int
	Gender    string
	Race      string
	Location  string
	Education string
	Language  string
	bracket   SocioEcon
}

// Socioeconomic status struct
type SocioEcon struct {
}

// NewDemographics is a factory for Demographic struct. !!Need to implement socioecon!!
func NewDemographics(age int, gender, race, location, education, language string) Demographics {
	// TODO: Enumeration checking logic
	return Demographics{age, gender, race, location, education, language, SocioEcon{}}
}
