package v1

import (
	"context"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// GetAll implements repository.TaskRepository
func (u *taskUseCaseV1) GetAll(ctx context.Context) (*[]model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return u.taskRepo.GetAll(ctx)
}

// GetByID implements repository.TaskRepository
func (u *taskUseCaseV1) GetByID(ctx context.Context, id uint) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return u.taskRepo.GetByID(ctx, id)
}
