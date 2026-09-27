package job

import "fmt"

type Service struct{
	jobs []Job
}

func NewService() *Service {
	return &Service{
	}
}

func (s *Service) CreateJob(id string, jobType string, payload string) Job{
	newJob := Job{
		ID: id,
		Type: jobType,
		Payload: payload,
		Status: "queued",
	}
	s.jobs = append(s.jobs, newJob)
	return newJob
}
func (s *Service) GetJobs() []Job{
	return s.jobs
}
func (s *Service) GetJob(id string) (Job, error){
	for _, job := range s.jobs {
		if job.ID == id{
			return job, nil
		}
	}
	return Job{}, fmt.Errorf("job with id %s not found", id)
}