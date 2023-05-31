package v1

import (
	"context"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Update implements repository.TaskRepository
func (u *taskUseCaseV1) Update(ctx context.Context, task *model.TaskUpdate) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	existingTask, err := u.taskRepo.GetByID(ctx, task.ID)
	if err != nil {
		return nil, err
	}

	if task.Title != "" {
		existingTask.Title = task.Title
	}

	if task.Description != "" {
		existingTask.Description = task.Description
	}

	if task.IsCompleted != nil {
		existingTask.IsCompleted = *task.IsCompleted
	}

	return u.taskRepo.Update(ctx, existingTask)
}
