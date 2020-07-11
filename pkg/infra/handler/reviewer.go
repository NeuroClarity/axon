package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

// "/api/reviewer/reviewJob" POST request body

type ReviewJobRequest struct {
	UID          string
	Webcam       bool
	Headset      RequestHeadset
	Demographics RequestDemographics
}

type RequestHeadset struct {
	Connected bool
	Type      string
}

type RequestDemographics struct {
	Age    int
	Gender string
	Race   string
}

// "/api/reviewer/reviewJob" POST response body

type ReviewJobResponse struct {
	StudyID int
	Content string
}

// ReviewerProfile retrieves profile information for a logged in Reviewer
func (rh *reviewerHandler) Ping(w http.ResponseWriter, r *http.Request) {
	ping := app.Ping()
	fmt.Printf("ping\n")
	fmt.Fprint(w, ping)
}

// AssignReviewJob retrieves a ReviewJob for a User based on Demographics and BioHardware criteria.
func (rh *reviewerHandler) AssignReviewJob(w http.ResponseWriter, r *http.Request) {
	var rjr ReviewJobRequest

	// Parsing logic and error handling in `parser.go`.
	err := decodeJSONBody(w, r, &rjr)
	if err != nil {
		log.Print(err.Error())

		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	fmt.Printf("%+v\n", rjr)

	reviewer, err := rh.reviewerRepo.GetReviewer(rjr.UID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	hardware, err := core.NewHardware(true, rjr.Headset.Type)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reviewJob, err := app.AssignReviewJob(reviewer, hardware, rh.reviewJobRepo)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Constructing response struct and marshaling into JSON.
	sid := reviewJob.Study.UID
	content := reviewJob.Study.Content.Location
	response := ReviewJobResponse{StudyID: sid, Content: content}

	js, err := json.Marshal(response)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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

			rawFirstName := user.(*jwt.Token).Claims.(jwt.MapClaims)["https://synapse.neuroclarity.ai/given_name"]
			firstName, ok := rawFirstName.(string)
			if !ok {
				log.Printf("Failed to get 'firstName' claim from JWT, got: %v.\n", rawFirstName)
				http.Error(w, "Incomplete firstName information in JWT", http.StatusBadRequest)
				return
			}

			rawLastName := user.(*jwt.Token).Claims.(jwt.MapClaims)["https://synapse.neuroclarity.ai/family_name"]
			lastName, ok := rawLastName.(string)
			if !ok {
				log.Printf("Failed to get 'lastName' claim from JWT, got: %v.\n", rawLastName)
				http.Error(w, "Incomplete lastName information in JWT", http.StatusBadRequest)
				return
			}

			rawEmail := user.(*jwt.Token).Claims.(jwt.MapClaims)["https://synapse.neuroclarity.ai/email"]
			email, ok := rawEmail.(string)
			if !ok {
				log.Printf("Failed to get 'email' claim from JWT, got: %v.\n", rawEmail)
				http.Error(w, "Incomplete email information in JWT", http.StatusBadRequest)
				return
			}

			demos := core.Demographics{}
			rh.reviewerRepo.NewReviewer(uid, firstName, lastName, email, demos)
		}
		next.ServeHTTP(w, r)
	})
}
