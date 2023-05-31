package postgres

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Create implements repository.TaskRepository
func (r *postgresTaskRepository) Create(ctx context.Context, task *model.Task) (uint, error) {
	res := r.Conn.WithContext(ctx).Create(task)
	if res.Error != nil {
		return 0, res.Error
	}

	return task.ID, nil
}
