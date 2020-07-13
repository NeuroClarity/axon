package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	app "github.com/NeuroClarity/axon/pkg/application"
	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/dgrijalva/jwt-go"
)

// ReviewerHandler deals with operations in the Reviewer context.
type ReviewerHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	AssignReviewJob(w http.ResponseWriter, r *http.Request)
	FinishReviewJob(w http.ResponseWriter, r *http.Request)
	CheckForReviewer(http.HandlerFunc) http.Handler
}

// reviewerHandler is internal implementation of ReviewerHandler.
type reviewerHandler struct {
	reviewerRepo     repo.ReviewerRepository
	reviewJobRepo    repo.ReviewJobRepository
	analyticsJobRepo repo.AnalyticsJobRepository
	studyRepo        repo.StudyRepository
}

// NewReviewerHandler is a factory for a ReviewerHandler.
func NewReviewerHandler(rr repo.ReviewerRepository, rjr repo.ReviewJobRepository, ajr repo.AnalyticsJobRepository, sr repo.StudyRepository) ReviewerHandler {
	return &reviewerHandler{rr, rjr, ajr, sr}
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

	reviewer, err := rh.reviewerRepo.GetReviewer(rjr.ReviewerID)
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
	content := reviewJob.Study.Content.VideoLocation
	response := ReviewJobResponse{StudyID: sid, Content: content}

	js, err := json.Marshal(response)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// FinishReviewJob turns a ReviewJob with Biometrics into an AnalyticsJob for processing into Insights.
func (rh *reviewerHandler) FinishReviewJob(w http.ResponseWriter, r *http.Request) {
	var req FinishReviewJobRequest

	// Parsing logic and error handling in `parser.go`.
	err := decodeJSONBody(w, r, &req)
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

	fmt.Printf("%+v\n", req)

	// Making domain objects to call our app logic.

	reviewer, err := rh.reviewerRepo.GetReviewer(req.Biometrics.ReviewerID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	study, err := rh.studyRepo.GetStudy(req.StudyID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	completed, err := time.Parse(time.RFC3339, req.Biometrics.Time)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	biometrics, err := core.NewBiometrics(reviewer, req.Biometrics.EEGData.Location, req.Biometrics.WebcamData.Location, completed)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Calling our app logic.

	if err := app.FinishReviewJob(reviewer, study, biometrics, rh.reviewJobRepo, rh.reviewerRepo, rh.analyticsJobRepo); err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Constructing response struct and marshaling into JSON.

	response := FinishReviewJobResponse{Success: true}

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

// "/api/reviewer/reviewJob" POST request body

type ReviewJobRequest struct {
	ReviewerID   string
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

// "/api/reviewer/finishReviewJob" POST request body

type FinishReviewJobRequest struct {
	StudyID    int
	Biometrics RequestBiometrics
}

type RequestBiometrics struct {
	ReviewerID string
	EEGData    RequestEEGData
	WebcamData WebcamData
	Time       string
}

type RequestEEGData struct {
	Location string
}

type WebcamData struct {
	Location string
}

// "/api/reviewer/finishReviewJob" POST response body

type FinishReviewJobResponse struct {
	Success bool
}
