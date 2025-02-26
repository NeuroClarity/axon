// Server instance for routing API endpoints.
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/infra/database"
	"github.com/NeuroClarity/axon/pkg/infra/handler"
	"github.com/NeuroClarity/axon/pkg/infra/middleware"
	"github.com/NeuroClarity/axon/pkg/infra/session"
	"github.com/NeuroClarity/axon/pkg/infra/storage"
	"github.com/rs/cors"
)

func main() {

	router := httprouter.New()

	// Service and middleware singletons.
	db, err := database.NewDatabase("neuroc", "NeuroCDB12", "nc-database.cr7v5oc2x2xe.us-west-1.rds.amazonaws.com", "5432", "postgres")
	if err != nil {
		log.Fatal(err.Error())
	}
	awsSession, err := session.NewSession("us-west-1")
	if err != nil {
		log.Fatal(err.Error())
	}
	s3, err := storage.NewStorage(awsSession.GetSession())
	if err != nil {
		log.Fatal(err.Error())
	}

	reviewRepo := repo.NewReviewerRepository(db)
	reviewJobRepo := repo.NewReviewJobRepository(db)
	analyticsJobRepo := repo.NewAnalyticsJobRepository(db)
	creatorRepo := repo.NewCreatorRepository(db)
	studyRepo := repo.NewStudyRepository(db, s3)
	jwtMiddleware := middleware.NewJWTMiddleware()

	// Dependency injection.
	publicHandler := handler.NewPublicHandler()
	reviewerHandler := handler.NewReviewerHandler(reviewRepo, reviewJobRepo, analyticsJobRepo, studyRepo)
	creatorHandler := handler.NewCreatorHandler(creatorRepo, studyRepo)

	// Public routes.
	router.HandlerFunc("GET", "/api/ping", publicHandler.Ping)

	// Reviewer routes.
	router.Handler("GET", "/api/reviewer/ping", jwtMiddleware.Handler(reviewerHandler.CheckForReviewer(http.HandlerFunc(reviewerHandler.Ping))))
	router.Handler("POST", "/api/reviewer/reviewJob", jwtMiddleware.Handler(reviewerHandler.CheckForReviewer(http.HandlerFunc(reviewerHandler.AssignReviewJob))))
	router.Handler("POST", "/api/reviewer/finishReviewJob", jwtMiddleware.Handler(reviewerHandler.CheckForReviewer(http.HandlerFunc(reviewerHandler.FinishReviewJob))))

	// Creator routes.
	router.Handler("GET", "/api/creator/ping", jwtMiddleware.Handler(creatorHandler.CheckForCreator(http.HandlerFunc(creatorHandler.Ping))))
	router.Handler("POST", "/api/creator/study", jwtMiddleware.Handler(creatorHandler.CheckForCreator(http.HandlerFunc(creatorHandler.CreateStudy))))
	router.Handler("POST", "/api/creator/results", jwtMiddleware.Handler(creatorHandler.CheckForCreator(http.HandlerFunc(creatorHandler.ViewStudy))))

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	http.ListenAndServe(":8000", corsWrapper.Handler(router))
}
