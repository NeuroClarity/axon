package app

import "fmt"

func ReviewerRegister() string {
	return "Reviewer Register. \n"
}

func ReviewerLogin(uid int) string {
	return fmt.Sprintf("Reviewer Login uid: %d.\n", uid)
}

func AssignReviewJob(uid int) string {
	return fmt.Sprintf("Assigning ReviewJob to %d. \n", uid)
}
