package core

import "time"

// Biometrics is any form of raw data generated from a Hardware.
type Biometrics struct {
	Reviewer   *Reviewer
	EEGData    EEGData
	WebcamData WebcamData
	Time       time.Time
}

// EEGData is raw EEG waveform data from a headset.
type EEGData struct {
	Location string
}

// WebcamData is raw webcam data.
type WebcamData struct {
	Location string
}

func NewBiometrics(reviewer *Reviewer, eegDataLocation, webcamDataLocation string, time time.Time) (*Biometrics, error) {
	webcam := WebcamData{webcamDataLocation}
	eeg := EEGData{eegDataLocation}
	return &Biometrics{reviewer, eeg, webcam, time}, nil
}
