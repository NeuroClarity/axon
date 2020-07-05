// Command line utility for axon; useful for development and quick testing.
package main

import (
	"fmt"

	"github.com/NeuroClarity/axon/pkg/domain/core"
	"github.com/NeuroClarity/axon/pkg/domain/factory"
	"github.com/NeuroClarity/axon/pkg/domain/repo"
	"github.com/NeuroClarity/axon/pkg/domain/services"
	"github.com/NeuroClarity/axon/pkg/infra/database"
)

func main() {
	// Test study functionality.
	database := database.NewDatabase("example.com", "SECRET")
	reviewJobRepo := repo.NewReviewJobRepository(database)
	reviewManager := services.NewReviewManager(reviewJobRepo)

	// Mock compiling a Study into ReviewJobs
	study := core.NewStudy(10, nil, nil, nil)
	reviewManager.MakeReviewsFrom(study)

	// Mock new reviewer creation
	demographics := core.NewDemographics(20, "male", "white", "berkeley", "bachelors", "english")
	reviewer, _ := factory.NewReviewer("kenny", demographics)

	// Mock retrieval of ReviewJob for user
	reviewJob := reviewManager.AssignReviewTo(reviewer)

	// Retrieving  a ReviewJob - ReviewManager
	fmt.Println("Axon.")
}

// mock handler for when reviewer registers w/ neuroclarity
func register() (*core.Reviewer, error) {
	database := database.NewDatabase("example.com", "SECRET")
	demographics := core.NewDemographics(20, "male", "white", "berkeley", "bachelors", "english")

	reviewer, err := factory.NewReviewer("kenny", demographics, database)
	if err != nil {
		return nil, err
	}

	return reviewer, nil
}

// mock handler for when user logs in
func login() {
	database := database.NewDatabase("example.com", "SECRET")
	reviewerRepo := repo.NewReviewerRepository(database)

	reviewer := reviewerRepo.GetReviewer(0)
}

// mock retrieving a review job
func getReviewJob() {
	database := database.NewDatabase("example.com", "SECRET")
	reviewerRepo := repo.NewReviewerRepository(database)
	rjRepo := repo.NewReviewJobRepository(database)

	kenny := reviewerRepo.GetReviewer(0)
	rjManager := factory.NewReviewManager(rjRepo)

	rj, err := rjManager.AssignReviewTo(kenny)
	if err != nil {
	}
}
