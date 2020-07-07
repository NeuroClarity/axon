package handler

import (
	"fmt"
	"net/http"

	"github.com/NeuroClarity/axon/pkg/domain/repo"
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
	clientRepo repo.ClientRepository
	studyRepo  repo.StudyRepository
}

// NewClientHandler is a factory for a ClientHandler.
func NewClientHandler(cr repo.ClientRepository, sr repo.StudyRepository) ClientHandler {
	return &clientHandler{cr, sr}
}

// ClientRegister handles registering a Client with the database.
func (ch *clientHandler) ClientRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Client Register. \n")
}

// ClientRegister handles retrieving Client information from the database.
func (ch *clientHandler) ClientLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Client Login uid: %s.\n", httprouter.ParamsFromContext(r.Context()).ByName("uid"))
}

func (ch *clientHandler) CreateStudy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Study Created. \n")
}

func (ch *clientHandler) ViewStudy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Study sid: %s.\n", httprouter.ParamsFromContext(r.Context()).ByName("sid"))
}
