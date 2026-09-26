package job

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