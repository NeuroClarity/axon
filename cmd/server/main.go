// Server instance for routing API endpoints.
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/infra/database"
	"github.com/NeuroClarity/axon/pkg/infra/handler"
	"github.com/NeuroClarity/axon/pkg/infra/middleware"
	"github.com/rs/cors"
)

func main() {

	router := httprouter.New()

	// Service and middleware singletons.
	db := database.NewDatabase("foo", "foo")
	reviewRepo := repo.NewReviewerRepository(db)
	reviewJobRepo := repo.NewReviewJobRepository(db)
	creatorRepo := repo.NewCreatorRepository(db)
	studyRepo := repo.NewStudyRepository(db)
	jwtMiddleware := middleware.NewJWTMiddleware()

	// Dependency injection.
	publicHandler := handler.NewPublicHandler()
	reviewerHandler := handler.NewReviewerHandler(reviewRepo, reviewJobRepo)
	creatorHandler := handler.NewCreatorHandler(creatorRepo, studyRepo)

	// Public routes.
	router.HandlerFunc("GET", "/api/ping", publicHandler.Ping)
	router.HandlerFunc("GET", "/api/permissions", publicHandler.Permissions)

	// Reviewer routes.
	router.Handler("GET", "/api/reviewer/ping", jwtMiddleware.Handler(http.HandlerFunc(publicHandler.Ping)))
	router.Handler("GET", "/api/reviewer/profile", jwtMiddleware.Handler(http.HandlerFunc(reviewerHandler.Profile)))

	// Creator routes.
	router.Handler("GET", "/api/study", jwtMiddleware.Handler(http.HandlerFunc(creatorHandler.CreateStudy)))
	router.Handler("GET", "/api/study/:sid", jwtMiddleware.Handler(http.HandlerFunc(creatorHandler.ViewStudy)))

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	http.ListenAndServe(":8000", corsWrapper.Handler(router))
}
