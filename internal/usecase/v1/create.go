package v1

import (
	"context"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Create implements repository.TaskRepository
func (u *taskUseCaseV1) Create(ctx context.Context, task *model.Task) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return u.taskRepo.Create(ctx, task)
}
