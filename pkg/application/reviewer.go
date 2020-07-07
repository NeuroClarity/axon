package app

import "fmt"

// RegisterReviewer creates a new Reviewer and persists.
func RegisterReviewer() string {
	// Reviewer factory
	//
	return "Reviewer Register. \n"
}

func ReviewerLogin(uid int) string {
	return fmt.Sprintf("Reviewer Login uid: %d.\n", uid)
}

func AssignReviewJob(uid int) string {
	return fmt.Sprintf("Assigning ReviewJob to %d. \n", uid)
}
