package main

import (
	"fmt"

	"GoGopherQueue-GGQ/internal/job"
)

func main() {

	service := job.NewService()
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
		return
	}
	

	fmt.Println(job)
}