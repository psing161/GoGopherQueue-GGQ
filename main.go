package main

import (
	"fmt"

	"GoGopherQueue-GGQ/internal/job"
)

func main() {
	newJob := job.Job{
		ID:      "job-123",
		Type:    "email",
		Payload: "send welcome email",
		Status:  "queued",
	}

	fmt.Println(newJob)
}