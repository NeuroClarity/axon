// Server instance for routing API endpoints.
package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"

	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/infra/auth"
	"github.com/NeuroClarity/axon/pkg/infra/database"
	"github.com/NeuroClarity/axon/pkg/infra/handler"
	"github.com/NeuroClarity/axon/pkg/infra/middleware"
)

func main() {

	gob.Register(map[string]interface{}{})
	router := httprouter.New()

	// Service singletons.
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatal(err.Error())
	}
	sessionStore := sessions.NewFilesystemStore("", []byte("FOO"))
	db := database.NewDatabase("foo", "foo")
	reviewRepo := repo.NewReviewerRepository(db)
	reviewJobRepo := repo.NewReviewJobRepository(db)
	clientRepo := repo.NewClientRepository(db)
	studyRepo := repo.NewStudyRepository(db)

	// Dependency injection.
	publicHandler := handler.NewPublicHandler()
	reviewerHandler := handler.NewReviewerHandler(reviewRepo, reviewJobRepo, authenticator, sessionStore)
	clientHandler := handler.NewClientHandler(clientRepo, studyRepo, authenticator, sessionStore)

	// Public routes.
	router.HandlerFunc("GET", "/api/ping", publicHandler.Ping)
	router.HandlerFunc("GET", "/api/permissions", publicHandler.Permissions)

	// Reviewer routes.
	router.HandlerFunc("GET", "/api/reviewer/callback", reviewerHandler.ReviewerCallback)
	router.HandlerFunc("GET", "/api/reviewer/login", reviewerHandler.ReviewerLogin)
	router.HandlerFunc("GET", "/api/reviewer/logout", reviewerHandler.ReviewerLogout)

	router.HandlerFunc("GET", "/api/reviewer/ping", middleware.Authenticate(publicHandler.Ping, sessionStore))
	router.HandlerFunc("GET", "/api/reviewer/profile", middleware.Authenticate(reviewerHandler.ReviewerProfile, sessionStore))
	router.HandlerFunc("GET", "/api/assign/:uid", middleware.Authenticate(reviewerHandler.AssignReviewJob, sessionStore))

	// Client routes.
	router.HandlerFunc("GET", "/api/client", clientHandler.ClientRegister)
	router.HandlerFunc("GET", "/api/client/:uid", clientHandler.ClientLogin)
	router.HandlerFunc("GET", "/api/study", clientHandler.CreateStudy)
	router.HandlerFunc("GET", "/api/study/:sid", clientHandler.ViewStudy)

	log.Fatal(http.ListenAndServe(":8000", router))
}
