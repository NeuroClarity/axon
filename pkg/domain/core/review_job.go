package core

// TODO: A ReviewJob is a Review that has yet to be completed.
type ReviewJob struct {
	Study        Study
	Demographics Demographics
}

func NewReviewJob(study Study, demographics Demographics) ReviewJob {
	return ReviewJob{study, demographics}
}
