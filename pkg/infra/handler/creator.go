package handler

import (
	"fmt"
	"net/http"
	"strconv"

	app "github.com/NeuroClarity/axon/pkg/application"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/julienschmidt/httprouter"
)

// CreatorHandler deals with operations in the Client context.
type CreatorHandler interface {
	CreateStudy(w http.ResponseWriter, r *http.Request)
	ViewStudy(w http.ResponseWriter, r *http.Request)
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

// CreateStudy handles creating and persisting a new Study.
func (ch *creatorHandler) CreateStudy(w http.ResponseWriter, r *http.Request) {
	study := app.CreateStudy()
	fmt.Fprint(w, study)
}

// CreateStudy handles viewing a persisted Study.
func (ch *creatorHandler) ViewStudy(w http.ResponseWriter, r *http.Request) {
	rawSID := httprouter.ParamsFromContext(r.Context()).ByName("sid")
	sid, err := strconv.Atoi(rawSID)
	if err != nil {
		// TODO
	}

	study := app.ViewStudy(sid)
	fmt.Fprint(w, study)
}
