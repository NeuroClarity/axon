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

	//"github.com/julienschmidt/httprouter"
	"github.com/dgrijalva/jwt-go"
)

// CreatorHandler deals with operations in the Client context.
type CreatorHandler interface {
	CreateStudy(w http.ResponseWriter, r *http.Request)
	ViewStudy(w http.ResponseWriter, r *http.Request)
	CheckForCreator(next http.HandlerFunc) http.Handler
}

// clientHandler is internal implementation of ClientHandler.
type creatorHandler struct {
	creatorRepo repo.CreatorRepository
	studyRepo   repo.StudyRepository
}

// NewClientHandler is a factory for a ClientHandler.
func NewCreatorHandler(cr repo.CreatorRepository, sr repo.StudyRepository) CreatorHandler {
	return &creatorHandler{cr, sr}
}

type CreateStudyRequest struct {
	UID         string
	VideoKey    string
	Demographic core.StudyRequest
}

type CreateStudyResponse struct {
	StudyID int
	URL     string
}

type ViewStudyRequest struct {
	StudyID int
}

type ViewStudyResponse struct {
}

// CreateStudy handles creating and persisting a new Study.
func (ch *creatorHandler) CreateStudy(w http.ResponseWriter, r *http.Request) {
	var csr CreateStudyRequest

	err := decodeJSONBody(w, r, &csr)
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
	fmt.Printf("%+v\n", csr)

	studyID, presignedUrl, err := app.CreateStudy(csr.UID, csr.VideoKey, &csr.Demographic, ch.studyRepo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	response, err := json.Marshal(CreateStudyResponse{
		StudyID: studyID,
		URL:     presignedUrl,
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// CreateStudy handles viewing a persisted Study.
func (ch *creatorHandler) ViewStudy(w http.ResponseWriter, r *http.Request) {
	var vsr ViewStudyRequest

	err := decodeJSONBody(w, r, &vsr)
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
	fmt.Printf("%+v\n", vsr)

	study, err := app.ViewStudy(vsr.StudyID, ch.studyRepo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	response, err := json.Marshal(study)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (ch *creatorHandler) CheckForCreator(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// JWT information from auth0 is now in the request context from JWTMiddleware.
		user := r.Context().Value("user")

		// Subject Claim is one of the values in the Payload of a JWT
		// (https://jwt.io/introduction/).  Here Auth0 is sending it back to
		// represent "what the JWT refers to" (eg. subject), and it contains
		// the Auth0 unique user_id. Hence, what we will use for DB UID.
		uid := user.(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)

		creator, err := ch.creatorRepo.GetCreator(uid)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else if creator != nil {
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

			company := ""
			ch.creatorRepo.NewCreator(uid, firstName, lastName, email, company)
		}
		next.ServeHTTP(w, r)
	})
}
