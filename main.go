package main

import (
	"fmt"

	"GoGopherQueue-GGQ/internal/job"
	"GoGopherQueue-GGQ/internal/repository"
)

func main() {
	// Create the in-memory repository
	repo := repository.NewMemoryRepository()

	// Inject the repository into the job service
	service := job.NewService(repo)

	// Create jobs
	_, err := service.CreateJob("job-123", "email", "Send welcome email")
	if err != nil {
		fmt.Println("Error creating job:", err)
		return
	}

	_, err = service.CreateJob("job-456", "report", "Generate monthly report")
	if err != nil {
		fmt.Println("Error creating job:", err)
		return
	}

	fmt.Println("Jobs created successfully")

	// Get a specific job
	foundJob, err := service.GetJob("job-123")
	if err != nil {
		fmt.Println("Error getting job:", err)
	} else {
		fmt.Printf("Found job: %+v\n", foundJob)
	}

	// Get all jobs
	fmt.Println("\nAll jobs:")
	for _, j := range service.GetJobs() {
		fmt.Printf("%+v\n", j)
	}

	// Delete a job
	err = service.DeleteJob("job-123")
	if err != nil {
		fmt.Println("Error deleting job:", err)
	} else {
		fmt.Println("\nJob job-123 deleted successfully")
	}

	// Display remaining jobs
	fmt.Println("\nRemaining jobs:")
	for _, j := range service.GetJobs() {
		fmt.Printf("%+v\n", j)
	}
}
