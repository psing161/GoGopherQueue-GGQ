package job

type Service struct {
	repo Repository

}

func NewService(repo Repository) *Service {
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

func (s *Service) GetJob(id string) (Job, error) {
	return s.repo.Get(id)
}

func (s *Service) DeleteJob(id string) error {
	return s.repo.Delete(id)
}