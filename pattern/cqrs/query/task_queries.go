package query

import (
	"cqrs/domain"
	"cqrs/repository"
)

type TaskQueryHandler struct {
	taskRepo repository.TaskRepository
}

func NewTaskQueryHander(taskRepo repository.TaskRepository) *TaskQueryHandler {
	return &TaskQueryHandler{taskRepo}
}

func (handler *TaskQueryHandler) HandleGetTaskByID(id string) (*domain.Task, error) {
	return handler.taskRepo.FindByID(id)
}

func (handler *TaskQueryHandler) HandleListTasks() ([]*domain.Task, error) {
	return handler.taskRepo.FindAll()
}
