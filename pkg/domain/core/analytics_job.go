package core

// AnalyticsJob is a collection of Biometrics collected from a Reviewer waiting to be processed into Insights.
type AnalyticsJob struct {
	// Email
	ReviewerID string
	// Locations in S3
	EEGData    string
	WebcamData string
}
