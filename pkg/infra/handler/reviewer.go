package handler

import (
	"fmt"
	"net/http"

	app "github.com/NeuroClarity/axon/pkg/application"
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/dgrijalva/jwt-go"
)

// ReviewerHandler deals with operations in the Reviewer context.
type ReviewerHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	AssignReviewJob(w http.ResponseWriter, r *http.Request)
	CheckForReviewer(http.HandlerFunc) http.Handler
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

// CheckForReviewer consults the ReviewerRepository to see if the Reviewer
// being referenced in the JWT (a unique auth0 id) exists already. If not, it
// will create persistent record before preceeding.
//
// This is a Middleware pattern.
func (rh *reviewerHandler) CheckForReviewer(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// JWT information from auth0 is now in the request context from JWTMiddleware.
		user := r.Context().Value("user")

		// Subject Claim is one of the values in the Payload of a JWT
		// (https://jwt.io/introduction/).  Here Auth0 is sending it back to
		// represent "what the JWT refers to" (eg. subject), and it contains
		// the Auth0 unique user_id. Hence, what we will use for DB UID.
		uid := user.(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)

		reviewer, err := rh.reviewerRepo.GetReviewer(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else if reviewer == nil {
			// Custom claims setup as auth0 rules from the client. These values
			// should always be in the JWT.
			firstName := user.(*jwt.Token).Claims.(jwt.MapClaims)["https://synapse.neuroclarity.ai/given_name"].(string)
			lastName := user.(*jwt.Token).Claims.(jwt.MapClaims)["https://synapse.neuroclarity.ai/family_name"].(string)
			email := user.(*jwt.Token).Claims.(jwt.MapClaims)["https://synapse.neuroclarity.ai/email"].(string)

			demos := core.Demographics{}
			rh.reviewerRepo.NewReviewer(uid, firstName, lastName, email, demos)
		}
		next.ServeHTTP(w, r)
	})
}
