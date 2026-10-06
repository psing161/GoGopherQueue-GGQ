package job


type Repository interface {
	Create(Job) error
	Get(id string) (Job, error)
	Delete(id string) error
	GetAll() []Job
}