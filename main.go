package main

import (
	"fmt"

	"GoGopherQueue-GGQ/internal/job"
	"GoGopherQueue-GGQ/internal/repository"
)

func main() {
	repo := repository.NewMemoryRepository()

	service := job.NewService(repo)

	service.CreateJob(
		"job-123",
		"email", 
		"send welcome email",
	)

	service.CreateJob(
		"job-456",
		"report", 
		"generate monthly report",
	)
	job, err := service.GetJob("job-999")

	if err != nil {
		fmt.Println("Error:", err)
	}else{
	fmt.Println(job)
	}

	deletedJob, err := service.DeleteJob("job-999")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Deleted:", deletedJob)
	

}