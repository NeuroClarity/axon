package core

// TODO: A Review is the response from a single user...
type Review struct {
	Reviewer Reviewer
	Insight  Insight
	Hardware Hardware
}

// TODO: NewReview is a factory for a finished Review.
func NewReview(reviewer Reviewer, insight Insight, hardware Hardware) {
}
