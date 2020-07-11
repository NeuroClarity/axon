package core

// Biometrics is any form of raw data generated from a Hardware.
type Biometrics struct {
	Reviewer   *Reviewer
	EEGData    EEGData
	WebcamData WebcamData
}

// EEGData is raw EEG waveform data from a headset.
type EEGData struct {
	S3Key string
}

// WebcamData is raw webcam data.
type WebcamData struct {
	S3Key string
}
