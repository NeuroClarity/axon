package core

type TargetAudience struct {
	NumParticipants int
	MinAge          int
	MaxAge          int
	Gender          string
	Race            string
	Eeg             bool
	EyeTracking     bool
}
