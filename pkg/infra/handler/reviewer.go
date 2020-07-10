package handler

import (
	"fmt"
	"net/http"

	app "github.com/NeuroClarity/axon/pkg/application"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
)

// ReviewerHandler deals with operations in the Reviewer context.
type ReviewerHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	AssignReviewJob(w http.ResponseWriter, r *http.Request)
}

// reviewerHandler is internal implementation of ReviewerHandler.
type reviewerHandler struct {
	reviewerRepo     repo.ReviewerRepository
	reviewJobRepo    repo.ReviewJobRepository
	analyticsJobRepo repo.AnalyticsJobRepository
}

// NewReviewerHandler is a factory for a ReviewerHandler.
func NewReviewerHandler(rr repo.ReviewerRepository, rjr repo.ReviewJobRepository, ajr repo.AnalyticsJobRepository) ReviewerHandler {
	return &reviewerHandler{rr, rjr, ajr}
}

// ReviewerProfile retrieves profile information for a logged in Reviewer
func (rh *reviewerHandler) Ping(w http.ResponseWriter, r *http.Request) {
	ping := app.Ping()
	fmt.Fprint(w, ping)
}

// AssignReviewJob retrieves a ReviewJob for a User based on Demographics and BioHardware criteria.
func (rh *reviewerHandler) AssignReviewJob(w http.ResponseWriter, r *http.Request) {
	// rawUID := httprouter.ParamsFromContext(r.Context()).ByName("uid")
	// uid, err := strconv.Atoi(rawUID)
	// if err != nil {
	// 	// TODO
	// }

	// reviewer, _ := app.AssignReviewJob(uid)
	// fmt.Fprint(w, reviewer)
}
