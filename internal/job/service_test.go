package job_test 

import (
	"testing"
	"GoGopherQueue-GGQ/internal/job" 
	"GoGopherQueue-GGQ/internal/repository" 
)

func TestServiceCreate(t *testing.T){
	repo := repository.NewMemoryRepository()

	service := job.NewService(repo)

	cjob , err := service.CreateJob("Job-123", "jobType", "payload")

	if err != nil{
		t.Fatalf("CreateJob() failed with error: %v", err)
	}
	if cjob.ID != "Job-123"{
		t.Errorf("Created Job Id does not match")
	}
}