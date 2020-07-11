package core

// TODO: A Review is the response from a single user...
type Review struct {
	Reviewer *Reviewer
	Insights Insights
	Hardware Hardware
}

// TODO: NewReview is a factory for a finished Review.
func NewReview(reviewer Reviewer, insight Insights, hardware Hardware) {
}
