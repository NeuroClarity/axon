package core

// TODO: A Reviewer is a person who participates in studies by watching videos and providing biometric data.
type Reviewer struct {
	UID          int
	Name         string
	Demographics Demographics
}

// TODO: Demographics enumerates characteristics about a Study Participant that are interesting to a Client.
type Demographics struct {
	Age       int
	Gender    string
	Race      string
	Location  string
	Education string
	Language  string
	bracket   SocioEcon
}

// TODO: Socioeconomic status struct
type SocioEcon struct {
}

// TODO: NewDemographics is a factory for Demographic struct. !!Need to implement socioecon!!
func NewDemographics(age int, gender, race, location, education, language string) Demographics {
	// TODO: Enumeration checking logic
	return Demographics{age, gender, race, location, education, language, SocioEcon{}}
}
