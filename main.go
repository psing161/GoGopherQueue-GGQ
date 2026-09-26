package main

import (
	"fmt"

	"GoGopherQueue-GGQ/internal/job"
)

func main() {

	service := job.NewService()
	newJob := service.CreateJob(
		"job-123",
		"email", 
		"send welcome email",
	)
	

	fmt.Println(newJob)
}