package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	app "github.com/NeuroClarity/axon/pkg/application"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/infra/auth"
	"github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

// ReviewerHandler deals with operations in the Reviewer context.
type ReviewerHandler interface {
	ReviewerCallback(w http.ResponseWriter, r *http.Request)
	ReviewerProfile(w http.ResponseWriter, r *http.Request)
	ReviewerLogin(w http.ResponseWriter, r *http.Request)
	ReviewerLogout(w http.ResponseWriter, r *http.Request)
	AssignReviewJob(w http.ResponseWriter, r *http.Request)
}

// reviewerHandler is internal implementation of ReviewerHandler.
type reviewerHandler struct {
	reviewRepo    repo.ReviewerRepository
	reviewJobRepo repo.ReviewJobRepository
	authenticator *auth.Authenticator
	sessionStore  *sessions.FilesystemStore
}

// NewReviewerHandler is a factory for a ReviewerHandler.
func NewReviewerHandler(rr repo.ReviewerRepository, rjr repo.ReviewJobRepository, auth *auth.Authenticator, session *sessions.FilesystemStore) ReviewerHandler {
	return &reviewerHandler{rr, rjr, auth, session}
}

// ReviewerCallback provides a callback route for Auth0.
func (rh *reviewerHandler) ReviewerCallback(w http.ResponseWriter, r *http.Request) {

	session, err := rh.sessionStore.Get(r, "auth-session")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		log.Print("Request URL state parameter does not match local session state.\n")
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	token, err := rh.authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Print("OAuth2 Token did not have id_token field (their service at fault).\n")
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: "ih5Ms51935CplO4inqwN1RL6mJxC5LMH",
	}

	idToken, err := rh.authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// User profile information retrieval.
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Registry with session profile: %+v\n", session.Values["profile"])
	http.Redirect(w, r, "/api/reviewer/profile", http.StatusSeeOther)
}

// ReviewerProfile retrieves profile information for a logged in Reviewer
func (rh *reviewerHandler) ReviewerProfile(w http.ResponseWriter, r *http.Request) {

	session, err := rh.sessionStore.Get(r, "auth-session")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Temp obviously.
	fmt.Fprintf(w, "user: %+v", session.Values["profile"])
}

// ReviewerLogin handles retrieving Reviewer information from the database.
func (rh *reviewerHandler) ReviewerLogin(w http.ResponseWriter, r *http.Request) {

	// Generates random state for encoding.
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := base64.StdEncoding.EncodeToString(b)
	session, err := rh.sessionStore.Get(r, "auth-session")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, rh.authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

// ReviewerLogout reroutes a Reviewer to securely logout through Auth0.
func (rh *reviewerHandler) ReviewerLogout(w http.ResponseWriter, r *http.Request) {

	domain := "dev-q7h0r088.us.auth0.com"

	logoutURL, err := url.Parse("https://" + domain)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Versioning developer specification from Auth0.
	logoutURL.Path += "/v2/logout"
	parameters := url.Values{}

	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", "ih5Ms51935CplO4inqwN1RL6mJxC5LMH")
	logoutURL.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}

// AssignReviewJob retrieves a ReviewJob for a User based on Demographics and BioHardware criteria.
func (rh *reviewerHandler) AssignReviewJob(w http.ResponseWriter, r *http.Request) {
	rawUID := httprouter.ParamsFromContext(r.Context()).ByName("uid")
	uid, err := strconv.Atoi(rawUID)
	if err != nil {
		// TODO
	}

	reviewer := app.AssignReviewJob(uid)
	fmt.Fprint(w, reviewer)
}
