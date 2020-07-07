package handler

import (
	"fmt"
	"net/http"
)

// PublicHandler routes non-authenticated API requests.
type PublicHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
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
	fmt.Fprint(w, "Success. \n")
}
