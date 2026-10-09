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

func TestServiceGetJob(t *testing.T) {
	repo := repository.NewMemoryRepository()
	service := job.NewService(repo)

	_, err := service.CreateJob("job-123", "email", "Welcome email")
	if err != nil {
		t.Fatalf("CreateJob() failed: %v", err)
	}

	got, err := service.GetJob("job-123")
	if err != nil {
		t.Fatalf("GetJob() failed: %v", err)
	}

	if got.ID != "job-123" {
		t.Errorf("GetJob() returned ID %q, want %q", got.ID, "job-123")
	}
}

func TestServiceGetMissingJob(t *testing.T) {
	repo := repository.NewMemoryRepository()
	service := job.NewService(repo)

	_, err := service.GetJob("missing-job")
	if err == nil {
		t.Fatal("GetJob() should return an error for a missing job")
	}
}

func TestServiceDeleteJob(t *testing.T) {
	repo := repository.NewMemoryRepository()
	service := job.NewService(repo)

	_, err := service.CreateJob("job-123", "email", "Welcome email")
	if err != nil {
		t.Fatalf("CreateJob() failed: %v", err)
	}

	if err := service.DeleteJob("job-123"); err != nil {
		t.Fatalf("DeleteJob() failed: %v", err)
	}

	_, err = service.GetJob("job-123")
	if err == nil {
		t.Fatal("GetJob() should return an error after deletion")
	}
}