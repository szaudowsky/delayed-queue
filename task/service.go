package task

type Service interface {
	Create(t Task) (Task, error)
	Tasks() ([]Task, error)
	Task(id string) (Task, error)
}

type service struct {
	repo Repository
}

func NewService(redisRepo Repository) Service {
	return &service{repo: redisRepo}
}

func (s *service) Create(t Task) (Task, error) {
	return s.repo.Create(t)
}

func (s *service) Tasks() ([]Task, error) {
	return s.repo.Tasks()
}

func (s *service) Task(id string) (Task, error) {
	return s.repo.Task(id)
}
