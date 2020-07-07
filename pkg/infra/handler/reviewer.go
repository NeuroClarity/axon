package handler

import (
	"fmt"
	"net/http"
	"strconv"

	app "github.com/NeuroClarity/axon/pkg/application"
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
	reviewer := app.ReviewerRegister()
	fmt.Fprint(w, reviewer)
}

// ReviewerLogin handles retrieving Reviewer information from the database.
func (rh *reviewerHandler) ReviewerLogin(w http.ResponseWriter, r *http.Request) {

	rawUID := httprouter.ParamsFromContext(r.Context()).ByName("uid")
	uid, err := strconv.Atoi(rawUID)
	if err != nil {
		// TODO
	}

	reviewer := app.ReviewerLogin(uid)
	fmt.Fprint(w, reviewer)
}

// AssignReviewJob retrieves a ReviewJob for a User based on Demographics and BioHardware criteria.
func (rh *reviewerHandler) AssignReviewJob(w http.ResponseWriter, r *http.Request) {
	rawUID := httprouter.ParamsFromContext(r.Context()).ByName("uid")
	uid, err := strconv.Atoi(rawUID)
	if err != nil {
		// TODO
	}

	reviewer := app.AssignReviewJob(uid)
	fmt.Fprint(w, reviewer)
}
