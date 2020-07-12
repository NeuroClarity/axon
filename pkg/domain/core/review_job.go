package core

import "time"

// A ReviewJob is a Review that needs to be completed.
type ReviewJob struct {
	Study     *Study
	Completed time.Time
}
