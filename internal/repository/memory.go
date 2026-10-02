package repository

import (
	"fmt"

	"GoGopherQueue-GGQ/internal/job"
)

type MemoryRepository struct {
	jobs []job.Job
}

func (r * MemoryRepository) Create(newJob Job.job) error{
	r.jobs = append(r.jobs, newJob)
	return nil

}

