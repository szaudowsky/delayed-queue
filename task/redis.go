package task

import (
	"github.com/go-redis/redis/v8"
)

type Repository interface {
	Create(t Task) (Task, error)
	Tasks() ([]Task, error)
	Task(id string) (Task, error)
}

type repository struct {
	client *redis.Client
}

func NewRepository(c *redis.Client) Repository {
	return &repository{client: c}
}

func (r *repository) Create(t Task) (Task, error) {
	return Task{}, nil
}

func (r *repository) Tasks() ([]Task, error) {
	tasks := make([]Task, 0)
	return tasks, nil
}

func (r *repository) Task(id string) (Task, error) {
	return Task{}, nil
}
