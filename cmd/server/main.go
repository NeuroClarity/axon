// Server instance for routing API endpoints.
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/infra/database"
	"github.com/NeuroClarity/axon/pkg/infra/handler"
)

func main() {

	router := httprouter.New()

	// Service singletons.
	db := database.NewDatabase("foo", "foo")
	reviewRepo := repo.NewReviewerRepository(db)
	reviewJobRepo := repo.NewReviewJobRepository(db)
	clientRepo := repo.NewClientRepository(db)
	studyRepo := repo.NewStudyRepository(db)

	// Dependency injection.
	publicHandler := handler.NewPublicHandler()
	reviewerHandler := handler.NewReviewerHandler(reviewRepo, reviewJobRepo)
	clientHandler := handler.NewClientHandler(clientRepo, studyRepo)

	// Public routes.
	router.HandlerFunc("GET", "/api/ping", publicHandler.Ping)

	// Reviewer routes.
	router.HandlerFunc("GET", "/api/reviewer", reviewerHandler.ReviewerRegister)
	router.HandlerFunc("GET", "/api/reviewer/:uid", reviewerHandler.ReviewerLogin)
	router.HandlerFunc("GET", "/api/assign/:uid", reviewerHandler.AssignReviewJob)

	// Client routes.
	router.HandlerFunc("GET", "/api/client", clientHandler.ClientRegister)
	router.HandlerFunc("GET", "/api/client/:uid", clientHandler.ClientLogin)
	router.HandlerFunc("GET", "/api/study", clientHandler.CreateStudy)
	router.HandlerFunc("GET", "/api/study/:sid", clientHandler.ViewStudy)

	log.Fatal(http.ListenAndServe(":8000", router))
}
