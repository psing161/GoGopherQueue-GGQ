package repository

import (
	"fmt"

	"GoGopherQueue-GGQ/internal/job"
)

type MemoryRepository struct {
	jobs []job.Job
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (r *MemoryRepository) Create(newJob job.Job) error {
	r.jobs = append(r.jobs, newJob)
	return nil
}

func (r *MemoryRepository) Get(id string) (job.Job, error) {
	for _, j := range r.jobs {
		if j.ID == id {
			return j, nil
		}
	}

	return job.Job{}, fmt.Errorf("job with id %s not found", id)
}

func (r *MemoryRepository) Delete(id string) error {
	for i, j := range r.jobs {
		if j.ID == id {
			r.jobs = append(r.jobs[:i], r.jobs[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("job with id %s not found", id)
}

func (r *MemoryRepository) GetAll() []job.Job {
	return r.jobs
}
