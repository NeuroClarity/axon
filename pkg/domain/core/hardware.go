package core

// Hardware
type Hardware struct {
	Webcam     Webcam
	EEGHeadset EEGHeadset
}

// EEGHeadset stores information about a Reviewer's eeg headset model for
// analytics.
type EEGHeadset struct {
	model string
}

// Webcam stores information about a Reviewer's webcam model if relevant for
// analytics.
type Webcam struct {
}

// NewHardware is a factory method for Hardware.
func NewHardware(webcam bool, eegModel string) (Hardware, error) {

	// TODO this is where we parse and organize logic about hardware specs
	// relevant to ml.

	return Hardware{Webcam{}, EEGHeadset{eegModel}}, nil
}
