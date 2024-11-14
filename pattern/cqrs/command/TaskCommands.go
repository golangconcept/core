package command

import (
	// "cqrs/domain"
	"cqrs/domain"
	"cqrs/repository"
)

type CreateTaskCommand struct {
	ID       string
	Title    string
	Status   string
	Assignee string
}

type UpdateTaskStatusCommand struct {
	ID     string
	Status string
}

type TaskCommandHandler struct {
	taskRepo repository.TaskRepository
}

func NewTaskCommandHandler(taskRepo repository.TaskRepository) *TaskCommandHandler {
	return &TaskCommandHandler{taskRepo}
}

func (handler *TaskCommandHandler) HandleCreateTask(command CreateTaskCommand) error {
	// Create task and save it to the repository
	task := domain.Task{
		ID:       command.ID,
		Title:    command.Title,
		Status:   command.Status,
		Assignee: command.Assignee,
	}
	return handler.taskRepo.Save(&task)
}

func (handler *TaskCommandHandler) HandleUpdateTaskStatus(command UpdateTaskStatusCommand) error {
	// Fetch the task, update its status, and save it
	task, err := handler.taskRepo.FindByID(command.ID)
	if err != nil {
		return err
	}

	task.Status = command.Status
	return handler.taskRepo.Save(task)
}
