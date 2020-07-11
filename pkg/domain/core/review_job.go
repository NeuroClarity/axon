package core

// TODO: A ReviewJob is a Review that has yet to be completed.
type ReviewJob struct {
	Study        Study
}

func NewReviewJob(study Study) ReviewJob {
	return ReviewJob{study}
}
