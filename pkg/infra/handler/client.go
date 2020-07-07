package handler

import (
	"fmt"
	"net/http"
	"strconv"

	app "github.com/NeuroClarity/axon/pkg/application"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/infra/auth"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

// ClientHandler deals with operations in the Client context.
type ClientHandler interface {
	ClientRegister(w http.ResponseWriter, r *http.Request)
	ClientLogin(w http.ResponseWriter, r *http.Request)
	CreateStudy(w http.ResponseWriter, r *http.Request)
	ViewStudy(w http.ResponseWriter, r *http.Request)
}

// clientHandler is internal implementation of ClientHandler.
type clientHandler struct {
	clientRepo    repo.ClientRepository
	studyRepo     repo.StudyRepository
	authenticator *auth.Authenticator
	sessionStore  *sessions.FilesystemStore
}

// NewClientHandler is a factory for a ClientHandler.
func NewClientHandler(cr repo.ClientRepository, sr repo.StudyRepository, auth *auth.Authenticator, session *sessions.FilesystemStore) ClientHandler {
	return &clientHandler{cr, sr, auth, session}
}

// ClientRegister handles registering a Client with the database.
func (ch *clientHandler) ClientRegister(w http.ResponseWriter, r *http.Request) {
	client := app.ClientRegister()
	fmt.Fprint(w, client)
}

// ClientRegister handles retrieving Client information from the database.
func (ch *clientHandler) ClientLogin(w http.ResponseWriter, r *http.Request) {

}

// CreateStudy handles creating and persisting a new Study.
func (ch *clientHandler) CreateStudy(w http.ResponseWriter, r *http.Request) {
	study := app.CreateStudy()
	fmt.Fprint(w, study)
}

// CreateStudy handles viewing a persisted Study.
func (ch *clientHandler) ViewStudy(w http.ResponseWriter, r *http.Request) {
	rawSID := httprouter.ParamsFromContext(r.Context()).ByName("sid")
	sid, err := strconv.Atoi(rawSID)
	if err != nil {
		// TODO
	}

	study := app.ViewStudy(sid)
	fmt.Fprint(w, study)
}
