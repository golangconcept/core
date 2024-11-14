package repository

import (
	"cqrs/domain"
	"errors"
)

type TaskRepository interface {
	Save(task *domain.Task) error
	FindByID(id string) (*domain.Task, error)
	FindAll() ([]*domain.Task, error)
}

type InMemoryTaskRepository struct {
	tasks map[string]*domain.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

func (repo *InMemoryTaskRepository) Save(task *domain.Task) error {
	repo.tasks[task.ID] = task
	return nil
}

func (repo *InMemoryTaskRepository) FindByID(id string) (*domain.Task, error) {
	task, exists := repo.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (repo *InMemoryTaskRepository) FindAll() ([]*domain.Task, error) {
	var tasks []*domain.Task
	for _, task := range repo.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
