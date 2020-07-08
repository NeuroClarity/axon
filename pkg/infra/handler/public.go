package handler

import (
	"fmt"
	"net/http"

	app "github.com/NeuroClarity/axon/pkg/application"
)

// PublicHandler routes non-authenticated API requests.
type PublicHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	Permissions(w http.ResponseWriter, r *http.Request)
}

// PublicHandler is internal implementation of PublicHandler.
type publicHandler struct {
}

// NewPublicHandler is a factory for a PublicHandler. publicHandler has no dependencies (for now).
func NewPublicHandler() PublicHandler {
	return &publicHandler{}
}

// Ping is a quick way to see if our API works.
func (ph *publicHandler) Ping(w http.ResponseWriter, r *http.Request) {
	ping := app.Ping()
	fmt.Fprint(w, ping)
}

// Permissions is a notification that user lacks authorization for this resource.
func (ph *publicHandler) Permissions(w http.ResponseWriter, r *http.Request) {
	ping := app.Permissions()
	fmt.Fprint(w, ping)
}
