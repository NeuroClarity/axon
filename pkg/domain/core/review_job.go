package core

// TODO: A ReviewJob is a Review that has yet to be completed.
type ReviewJob struct {
	Study        Study
	Demographics Demographics
}

func NewReviewJob(study Study, demographics Demographics) ReviewJob {
	return ReviewJob{study, demographics}
}

//TODO: Calculate a priority score based on attributes for user distribution.
func Priority() int {
	return 1
}

// TODO: Take raw biometrics and generate insights, load into Review.
func Complete(biometric BioMetric) *Review {
	// generate insight from biometric
	// Make Review from insight, reviewer, hardware
}
