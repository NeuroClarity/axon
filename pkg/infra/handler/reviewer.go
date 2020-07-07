package handler

import (
	"fmt"
	"net/http"

	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/julienschmidt/httprouter"
)

// ReviewerHandler deals with operations in the Reviewer context.
type ReviewerHandler interface {
	ReviewerRegister(w http.ResponseWriter, r *http.Request)
	ReviewerLogin(w http.ResponseWriter, r *http.Request)
	AssignReviewJob(w http.ResponseWriter, r *http.Request)
}

// reviewerHandler is internal implementation of ReviewerHandler.
type reviewerHandler struct {
	reviewRepo    repo.ReviewerRepository
	reviewJobRepo repo.ReviewJobRepository
}

// NewReviewerHandler is a factory for a ReviewerHandler.
func NewReviewerHandler(rr repo.ReviewerRepository, rjr repo.ReviewJobRepository) ReviewerHandler {
	return &reviewerHandler{rr, rjr}
}

// ReviewerRegister handles registering a Reviewer with the database.
func (rh *reviewerHandler) ReviewerRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Reviewer Register. \n")
}

// ReviewerLogin handles retrieving Reviewer information from the database.
func (rh *reviewerHandler) ReviewerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reviewer Login uid: %s.\n", httprouter.ParamsFromContext(r.Context()).ByName("uid"))
}

func (rh *reviewerHandler) AssignReviewJob(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Assigning ReviewJob to %s. \n", httprouter.ParamsFromContext(r.Context()).ByName("uid"))
}
