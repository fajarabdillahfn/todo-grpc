package v1

import (
	"context"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Update implements repository.TaskRepository
func (u *taskUseCaseV1) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := u.taskRepo.GetByID(ctx, task.ID)
	if err != nil {
		return nil, err
	}

	return u.taskRepo.Update(ctx, task)
}
