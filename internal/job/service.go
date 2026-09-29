package job

import "fmt"
import "GoGopherQueue-GGQ/internal/repository"

type Service struct{
	repo repository.Repository

}

func NewService(repo repository.Repository) *Service {
    return &Service{
        repo: repo,
    }
}

func (s *Service) CreateJob(id string, jobType string, payload string) (Job, error) {
	newJob := Job{
		ID:      id,
		Type:    jobType,
		Payload: payload,
		Status:  "queued",
	}

	err := s.repo.Create(newJob)
	if err != nil {
		return Job{}, err
	}

	return newJob, nil
}
func (s *Service) GetJobs() []Job {
	return s.repo.GetAll()
}
func (s *Service) GetJob(id string) (Job, error){
	for _, job := range s.jobs {
		if job.ID == id{
			return job, nil
		}
	}
	return Job{}, fmt.Errorf("job with id %s not found", id)
}

func (s *Service) DeleteJob(id string)(Job, error){
	for i, job := range s.jobs{
		if job.ID == id{
			s.jobs = append(s.jobs[:i],s.jobs[i+1:]...)
			return job,nil
		}
	}
	return Job{},fmt.Errorf("job with id %s not found",id)

}